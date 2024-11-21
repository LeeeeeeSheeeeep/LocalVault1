// LocalVault App Logic

const API_BASE = 'http://localhost:8080/api';

// Web Audio API State
let audioCtx = null;
let rainNode = null;
let windNode = null;
let isPlayingAmbient = false;
let currentCondition = 'Clear';

// Webcam Scanner State
let webcamStream = null;
let scannerInterval = null;
let scannedSegments = {};

// Keystroke dynamics biometric variables
let keydownTimestamps = {};
let lastKeyReleaseTime = null;
let currentDwellTimes = [];
let currentFlightTimes = [];
let enrollAttemptsLeft = 3;
let enrolledEntries = [];

function resetKeystrokeTracking() {
    keydownTimestamps = {};
    lastKeyReleaseTime = null;
    currentDwellTimes = [];
    currentFlightTimes = [];
}

function handlePasswordKeyDown(e) {
    if (e.key === 'Enter') return;
    if (e.repeat) return;
    
    const now = performance.now();
    if (lastKeyReleaseTime !== null) {
        currentFlightTimes.push(now - lastKeyReleaseTime);
    }
    keydownTimestamps[e.key] = now;
}

function handlePasswordKeyUp(e) {
    if (e.key === 'Enter') return;
    
    const now = performance.now();
    const pressTime = keydownTimestamps[e.key];
    if (pressTime !== undefined) {
        currentDwellTimes.push(now - pressTime);
        delete keydownTimestamps[e.key];
    }
    lastKeyReleaseTime = now;
}

document.addEventListener('DOMContentLoaded', () => {
    // Navigation
    const navItems = document.querySelectorAll('.nav-item');
    const views = document.querySelectorAll('.view');

    navItems.forEach(item => {
        item.addEventListener('click', (e) => {
            e.preventDefault();
            
            // Handle active state
            navItems.forEach(nav => nav.classList.remove('active'));
            item.classList.add('active');

            // Switch views
            const targetViewId = 'view-' + item.dataset.view;
            views.forEach(view => {
                if (view.id === targetViewId) {
                    view.classList.remove('hidden');
                    view.classList.add('active');
                    
                    // Specific logic for Nexus view
                    if (item.dataset.view === 'nexus' && window.nexusApp) {
                        window.nexusApp.resize();
                        loadGraphData();
                    }

                    // Specific logic for Deadman view
                    if (item.dataset.view === 'deadman') {
                        loadDeadManStatus();
                    }
                } else {
                    view.classList.remove('active');
                    view.classList.add('hidden');
                }
            });
        });
    });

    // Nexus Init
    if (typeof NeuralNexus !== 'undefined') {
        window.nexusApp = new NeuralNexus('nexus-canvas');
        document.getElementById('btn-load-graph').addEventListener('click', loadGraphData);
    }

    async function loadGraphData() {
        if (!window.nexusApp) return;
        try {
            const res = await fetch(`${API_BASE}/graph`);
            const data = await res.json();
            if (data && data.nodes) {
                window.nexusApp.stop();
                window.nexusApp.loadData(data);
                window.nexusApp.start();
            }
        } catch(e) {
            console.error("Failed to load graph data", e);
        }
    }

    // Keystroke dynamics timing listeners
    const unlockInput = document.getElementById('unlock-password');
    if (unlockInput) {
        unlockInput.addEventListener('keydown', (e) => {
            if (e.key === 'Backspace') {
                unlockInput.value = '';
                resetKeystrokeTracking();
                return;
            }
            if (e.key === 'Enter') {
                attemptUnlock();
                return;
            }
            handlePasswordKeyDown(e);
        });
        unlockInput.addEventListener('keyup', handlePasswordKeyUp);
        unlockInput.addEventListener('focus', () => {
            unlockInput.value = '';
            resetKeystrokeTracking();
        });
    }

    const enrollInput = document.getElementById('enroll-keystroke-input');
    if (enrollInput) {
        enrollInput.addEventListener('keydown', (e) => {
            if (e.key === 'Backspace') {
                enrollInput.value = '';
                resetKeystrokeTracking();
                return;
            }
            if (e.key === 'Enter') {
                handleEnrollSubmit();
                return;
            }
            handlePasswordKeyDown(e);
        });
        enrollInput.addEventListener('keyup', handlePasswordKeyUp);
        enrollInput.addEventListener('focus', () => {
            enrollInput.value = '';
            resetKeystrokeTracking();
        });
    }

    // Search Logic
    const searchInput = document.getElementById('search-input');
    const searchResults = document.getElementById('search-results');
    const searchLoading = document.getElementById('search-loading');

    let debounceTimer;

    searchInput.addEventListener('input', (e) => {
        clearTimeout(debounceTimer);
        const query = e.target.value.trim();

        if (query.length === 0) {
            searchResults.innerHTML = `
                <div class="empty-state">
                    <svg viewBox="0 0 24 24" width="48" height="48" stroke="#ccc" stroke-width="1" fill="none"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
                    <h3>Ready to search</h3>
                    <p>Enter keywords to find anything in your personal vault.</p>
                </div>
            `;
            return;
		}

        debounceTimer = setTimeout(() => {
            performSearch(query);
        }, 300);
    });

    async function performSearch(query) {
        searchLoading.classList.remove('hidden');
        searchResults.innerHTML = '';

        try {
            const response = await fetch(`${API_BASE}/search?q=${encodeURIComponent(query)}&limit=20`);
            const data = await response.json();

            searchLoading.classList.add('hidden');

            if (!data || data.length === 0) {
                searchResults.innerHTML = `
                    <div class="empty-state">
                        <h3>No results found</h3>
                        <p>We couldn't find anything matching "${query}"</p>
                    </div>
                `;
                return;
            }

            // Render results
            data.forEach(doc => {
                const card = document.createElement('div');
                card.className = 'result-card';
                card.innerHTML = `
                    <div class="result-meta">
                        <span class="type-badge">${doc.DocType}</span>
                        <span>${doc.Author || 'Unknown'}</span>
                    </div>
                    <div class="result-title">${doc.Title}</div>
                    <div class="result-snippet">${escapeHTML(doc.Content)}</div>
                `;
                
                card.addEventListener('click', () => {
                    // Fetch weather for this document's creation date
                    fetchWeatherForDoc(doc.CreatedAt || new Date().toISOString());
                    alert(`Selected: ${doc.Title}\nTime: ${doc.CreatedAt || 'Unknown'}\nWeather ambient synced.`);
                });

                searchResults.appendChild(card);
            });
        } catch (error) {
            searchLoading.classList.add('hidden');
            searchResults.innerHTML = `
                <div class="empty-state">
                    <h3 style="color: red;">Search Failed</h3>
                    <p>Is the local backend running?</p>
                </div>
            `;
        }
    }
});

// Weather Ambient Fetch
async function fetchWeatherForDoc(dateStr) {
    try {
        const res = await fetch(`${API_BASE}/weather?date=${encodeURIComponent(dateStr)}`);
        const data = await res.json();
        
        document.getElementById('weather-temp').textContent = `${Math.round(data.temp)}°C`;
        document.getElementById('weather-desc').textContent = `Ambient ${data.condition}`;
        
        const iconEl = document.getElementById('weather-icon');
        if (data.condition === 'Rainy') iconEl.textContent = '🌧️';
        else if (data.condition === 'Windy') iconEl.textContent = '💨';
        else iconEl.textContent = '☀️';

        currentCondition = data.condition;
        
        if (isPlayingAmbient) {
            // Re-trigger audio synthesis immediately to morph the soundscapes
            startAmbientAudio(currentCondition);
        }
    } catch(e) {
        console.error("Failed to fetch weather data", e);
    }
}

// Web Audio Ambient Synthesis
function createNoiseBuffer(color, duration, sampleRate) {
    const bufferSize = sampleRate * duration;
    const buffer = audioCtx.createBuffer(1, bufferSize, sampleRate);
    const data = buffer.getChannelData(0);
    let lastOut = 0.0;
    
    for (let i = 0; i < bufferSize; i++) {
        const white = Math.random() * 2 - 1;
        if (color === 'brown') {
            data[i] = (lastOut + (0.02 * white)) / 1.02;
            lastOut = data[i];
            data[i] *= 3.5;
        } else {
            data[i] = white;
        }
    }
    return buffer;
}

function stopAmbientAudioNodes() {
    if (rainNode) {
        try { rainNode.source.stop(); } catch(e){}
        rainNode = null;
    }
    if (windNode) {
        try { 
            windNode.source.stop(); 
            windNode.lfo.stop();
        } catch(e){}
        windNode = null;
    }
}

function startAmbientAudio(condition) {
    if (!audioCtx) {
        audioCtx = new (window.AudioContext || window.webkitAudioContext)();
    }
    if (audioCtx.state === 'suspended') {
        audioCtx.resume();
    }

    stopAmbientAudioNodes();

    if (condition === 'Rainy') {
        const rainSource = audioCtx.createBufferSource();
        rainSource.buffer = createNoiseBuffer('brown', 2, audioCtx.sampleRate);
        rainSource.loop = true;

        const filter = audioCtx.createBiquadFilter();
        filter.type = 'lowpass';
        filter.frequency.value = 800;

        const gain = audioCtx.createGain();
        gain.gain.value = 0.25;

        rainSource.connect(filter);
        filter.connect(gain);
        gain.connect(audioCtx.destination);

        rainSource.start();
        rainNode = { source: rainSource, gain: gain };
    } else if (condition === 'Windy') {
        const windSource = audioCtx.createBufferSource();
        windSource.buffer = createNoiseBuffer('white', 2, audioCtx.sampleRate);
        windSource.loop = true;

        const filter = audioCtx.createBiquadFilter();
        filter.type = 'bandpass';
        filter.Q.value = 8;
        filter.frequency.value = 400;

        const lfo = audioCtx.createOscillator();
        lfo.frequency.value = 0.15;

        const lfoGain = audioCtx.createGain();
        lfoGain.gain.value = 250;

        const gain = audioCtx.createGain();
        gain.gain.value = 0.12;

        lfo.connect(lfoGain);
        lfoGain.connect(filter.frequency);
        windSource.connect(filter);
        filter.connect(gain);
        gain.connect(audioCtx.destination);

        lfo.start();
        windSource.start();
        windNode = { source: windSource, lfo: lfo, gain: gain };
    } else {
        // Clear: Gentle background breeze
        const windSource = audioCtx.createBufferSource();
        windSource.buffer = createNoiseBuffer('brown', 2, audioCtx.sampleRate);
        windSource.loop = true;

        const filter = audioCtx.createBiquadFilter();
        filter.type = 'lowpass';
        filter.frequency.value = 200;

        const gain = audioCtx.createGain();
        gain.gain.value = 0.05;

        windSource.connect(filter);
        filter.connect(gain);
        gain.connect(audioCtx.destination);

        windSource.start();
        rainNode = { source: windSource, gain: gain };
    }
}

function toggleAmbientAudio() {
    const btn = document.getElementById('btn-ambient-audio');
    if (isPlayingAmbient) {
        stopAmbientAudioNodes();
        isPlayingAmbient = false;
        btn.textContent = '🔊 Play Ambient';
        btn.classList.remove('playing');
    } else {
        startAmbientAudio(currentCondition);
        isPlayingAmbient = true;
        btn.textContent = '🔇 Mute Ambient';
        btn.classList.add('playing');
    }
}

// Dead Man's Switch Settings
async function loadDeadManStatus() {
    try {
        const res = await fetch(`${API_BASE}/deadman`);
        const data = await res.json();
        
        document.getElementById('dm-threshold').value = data.threshold_days;
        document.getElementById('dm-last-h').textContent = new Date(data.last_heartbeat).toLocaleString();
        document.getElementById('dm-state').textContent = data.triggered ? 'Triggered (EMERGENCY RELEASE)' : 'Active (Secured)';
        document.getElementById('dm-state').style.color = data.triggered ? 'red' : 'green';
        
        const remDays = data.triggered ? 0 : Math.max(0, data.threshold_days - Math.floor((new Date() - new Date(data.last_heartbeat)) / (1000 * 60 * 60 * 24)));
        document.getElementById('dm-rem-time').textContent = `${remDays} days remaining`;
    } catch(e) {
        console.error("Failed to load Dead Man Switch status", e);
    }
}

async function saveDeadManSettings() {
    const threshold = parseInt(document.getElementById('dm-threshold').value);
    const total = parseInt(document.getElementById('dm-total-shares').value);
    const min = parseInt(document.getElementById('dm-min-shares').value);
    const recipients = document.getElementById('dm-recipients').value.split('\n').map(e => e.trim()).filter(Boolean);

    try {
        const res = await fetch(`${API_BASE}/deadman`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                action: 'settings',
                threshold: threshold,
                total: total,
                min: min,
                recipients: recipients
            })
        });
        if (res.ok) {
            alert('Dead Man Switch settings saved successfully.');
            loadDeadManStatus();
        } else {
            const err = await res.text();
            alert(`Error: ${err}`);
        }
    } catch(e) {
        alert(`Failed to save: ${e.message}`);
    }
}

async function triggerHeartbeat() {
    try {
        const res = await fetch(`${API_BASE}/deadman`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ action: 'heartbeat' })
        });
        if (res.ok) {
            alert('Heartbeat sent. Vault longevity extended.');
            loadDeadManStatus();
        }
    } catch(e) {
        alert(`Heartbeat failed: ${e.message}`);
    }
}

// QR Paper Backup Generate
async function generateQRBackup() {
    const container = document.getElementById('qr-display-container');
    container.innerHTML = 'Generating cold-storage QR matrices...';

    try {
        const res = await fetch(`${API_BASE}/backup`);
        const data = await res.json();
        
        if (data.status === 'success' && data.segments) {
            container.innerHTML = '';
            data.segments.forEach((seg, idx) => {
                const segCard = document.createElement('div');
                segCard.style.display = 'flex';
                segCard.style.flexDirection = 'column';
                segCard.style.alignItems = 'center';
                segCard.style.background = '#f9f9f9';
                segCard.style.padding = '1rem';
                segCard.style.borderRadius = '8px';
                
                const canvas = document.createElement('canvas');
                segCard.appendChild(canvas);

                const label = document.createElement('div');
                label.style.fontSize = '0.8rem';
                label.style.fontWeight = '600';
                label.style.marginTop = '0.5rem';
                label.textContent = `Sheet ${seg.i} of ${seg.t} (ID: ${seg.c.toString(16)})`;
                segCard.appendChild(label);

                container.appendChild(segCard);

                // Render QR code onto canvas using QRious
                new QRious({
                    element: canvas,
                    value: JSON.stringify(seg),
                    size: 200,
                    level: 'M'
                });
            });
        } else {
            container.innerHTML = '<span style="color:red;">Failed to create QR backup.</span>';
        }
    } catch(e) {
        container.innerHTML = `<span style="color:red;">Error: ${e.message}</span>`;
    }
}

// Camera WebRTC QR Scanner using jsQR
async function toggleWebcam() {
    const video = document.getElementById('webcam-preview');
    const overlay = document.getElementById('webcam-overlay');
    const placeholder = document.getElementById('webcam-placeholder');
    const btn = document.getElementById('btn-toggle-camera');

    if (webcamStream) {
        // Stop camera
        clearInterval(scannerInterval);
        webcamStream.getTracks().forEach(track => track.stop());
        webcamStream = null;
        video.style.display = 'none';
        placeholder.style.display = 'flex';
        btn.textContent = 'Start Scanner';
        document.getElementById('scan-progress-box').classList.add('hidden');
    } else {
        // Start camera
        try {
            webcamStream = await navigator.mediaDevices.getUserMedia({ video: { facingMode: "environment" } });
            video.srcObject = webcamStream;
            video.setAttribute("playsinline", true);
            video.play();
            video.style.display = 'block';
            placeholder.style.display = 'none';
            btn.textContent = 'Stop Scanner';

            scannedSegments = {};
            document.getElementById('scan-progress-box').classList.remove('hidden');
            document.getElementById('scanned-count').textContent = '0';
            document.getElementById('scanned-total').textContent = '?';

            // Start QR scan loop
            const canvasCtx = overlay.getContext('2d');
            scannerInterval = setInterval(() => {
                if (video.readyState === video.HAVE_ENOUGH_DATA) {
                    overlay.width = video.videoWidth;
                    overlay.height = video.videoHeight;
                    canvasCtx.drawImage(video, 0, 0, overlay.width, overlay.height);
                    
                    const imgData = canvasCtx.getImageData(0, 0, overlay.width, overlay.height);
                    const code = jsQR(imgData.data, imgData.width, imgData.height, {
                        inversionAttempts: "dontInvert",
                    });

                    if (code) {
                        drawQRBoundingBox(code.location, canvasCtx);
                        try {
                            const segment = JSON.parse(code.data);
                            if (segment && segment.i && segment.t && segment.p) {
                                if (!scannedSegments[segment.i]) {
                                    scannedSegments[segment.i] = segment;
                                    const count = Object.keys(scannedSegments).length;
                                    document.getElementById('scanned-count').textContent = count;
                                    document.getElementById('scanned-total').textContent = segment.t;
                                    
                                    // Trigger sound effect/visual signal
                                    canvasCtx.fillStyle = "rgba(52, 199, 89, 0.4)";
                                    canvasCtx.fillRect(0, 0, overlay.width, overlay.height);

                                    if (count === segment.t) {
                                        // All segments gathered! Perform database recovery
                                        clearInterval(scannerInterval);
                                        restoreDatabaseFromScanned();
                                    }
                                }
                            }
                        } catch(e) {
                            // Ignored: Not a valid segment JSON
                        }
                    }
                }
            }, 150);
        } catch(e) {
            alert(`Unable to open camera: ${e.message}`);
        }
    }
}

function drawQRBoundingBox(location, ctx) {
    ctx.beginPath();
    ctx.moveTo(location.topLeftCorner.x, location.topLeftCorner.y);
    ctx.lineTo(location.topRightCorner.x, location.topRightCorner.y);
    ctx.lineTo(location.bottomRightCorner.x, location.bottomRightCorner.y);
    ctx.lineTo(location.bottomLeftCorner.x, location.bottomLeftCorner.y);
    ctx.closePath();
    ctx.lineWidth = 4;
    ctx.strokeStyle = "#34c759";
    ctx.stroke();
}

async function restoreDatabaseFromScanned() {
    const list = Object.values(scannedSegments);
    try {
        const res = await fetch(`${API_BASE}/backup`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(list)
        });
        const data = await res.json();
        if (res.ok && data.status === 'success') {
            alert(`Emergency Restore Successful!\nRebuilt and restored ${data.restored} documents.`);
            toggleWebcam(); // Close camera
        } else {
            alert(`Restore Failed: ${data.error || 'Invalid segments'}`);
        }
    } catch(e) {
        alert(`Restore API request failed: ${e.message}`);
    }
}

// Sync Connector Logic (Global accessible)
async function syncConnector(providerId) {
    const statusBadge = document.getElementById(`${providerId}-status`);
    const originalText = statusBadge.textContent;
    
    let config = {};
    if (providerId === 'github') {
        const token = document.getElementById('github-token').value;
        if (!token) {
            alert('Please enter a Personal Access Token');
            return;
        }
        config.token = token;
    } else if (providerId === 'rss') {
        const url = document.getElementById('rss-url').value;
        if (!url) {
            alert('Please enter an RSS URL');
            return;
        }
        config.url = url;
    }

    statusBadge.textContent = 'Syncing...';
    statusBadge.style.background = '#0071e3';
    statusBadge.style.color = '#fff';

    try {
        const response = await fetch(`${API_BASE}/sync`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                provider_id: providerId,
                config: config
            })
        });

        const data = await response.json();
        if (response.ok) {
            statusBadge.textContent = `Synced ${data.synced} items`;
            statusBadge.style.background = '#34c759';
            setTimeout(() => {
                statusBadge.textContent = 'Ready';
                statusBadge.style.background = '#e5e5ea';
                statusBadge.style.color = 'var(--text-muted)';
            }, 3000);
        } else {
            throw new Error(data.error || 'Failed to sync');
        }
    } catch (error) {
        console.error("Sync error:", error);
        statusBadge.textContent = 'Error';
        statusBadge.style.background = 'red';
    }
}

// Helper
function escapeHTML(str) {
    if (!str) return '';
    const div = document.createElement('div');
    div.textContent = str;
    return div.innerHTML;
}

// Deep Security: Dynamic Unlock
async function attemptUnlock() {
    const passwordInput = document.getElementById('unlock-password');
    const password = passwordInput.value;
    const errorEl = document.getElementById('unlock-error');
    
    if (!password) {
        alert('Please enter your vault password.');
        return;
    }

    const payload = {
        password: password,
        keystroke: {
            dwell_times: currentDwellTimes,
            flight_times: currentFlightTimes
        }
    };

    resetKeystrokeTracking();
    errorEl.classList.add('hidden');

    try {
        const res = await fetch(`${API_BASE}/unlock`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        });

        const data = await res.json();
        if (res.ok && data.status === 'success') {
            // Remove blur and hide overlay
            document.querySelector('.app-container').classList.remove('blur');
            document.getElementById('unlock-overlay').classList.add('fade-out');
            
            if (data.vault === 'decoy') {
                alert('Vault decrypted successfully.\n🔒 Warning: Decoy database active (rhythm signature mismatch or duress mode).');
            } else {
                alert('Vault decrypted successfully. Base data integrity verified.');
            }
        } else {
            errorEl.textContent = data.reason === 'invalid_credentials' ? 'Invalid credentials' : 'Authentication failed';
            errorEl.classList.remove('hidden');
            passwordInput.value = '';
        }
    } catch(e) {
        alert(`Unlock request failed: ${e.message}`);
    }
}

// Deep Security: Keystroke Enrollment
async function handleEnrollSubmit() {
    const enrollInput = document.getElementById('enroll-keystroke-input');
    const password = enrollInput.value;
    const attemptsLeftEl = document.getElementById('enroll-attempts-left');
    const fillEl = document.getElementById('enroll-progress-fill');
    const statusEl = document.getElementById('enroll-status');

    if (!password) {
        alert('Please type the enrollment password.');
        return;
    }

    // Save timing entry
    const entry = {
        dwell_times: currentDwellTimes,
        flight_times: currentFlightTimes
    };

    enrolledEntries.push(entry);
    resetKeystrokeTracking();
    enrollAttemptsLeft--;

    // Update progress UI
    attemptsLeftEl.textContent = enrollAttemptsLeft;
    const pct = ((3 - enrollAttemptsLeft) / 3) * 100;
    fillEl.style.width = `${pct}%`;
    statusEl.textContent = `Attempt ${3 - enrollAttemptsLeft} captured.`;

    enrollInput.value = '';

    if (enrollAttemptsLeft === 0) {
        statusEl.textContent = 'Submitting biometric timing signatures...';
        try {
            const res = await fetch(`${API_BASE}/keystroke/enroll`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ entries: enrolledEntries })
            });

            if (res.ok) {
                statusEl.textContent = 'Biometric profiling enrolled successfully! Keyboard biometrics active.';
                statusEl.style.color = '#34c759';
            } else {
                const text = await res.text();
                statusEl.textContent = `Enrollment failed: ${text}`;
                statusEl.style.color = 'red';
            }
        } catch(e) {
            statusEl.textContent = `Error: ${e.message}`;
            statusEl.style.color = 'red';
        }
        
        // Reset states after complete
        enrolledEntries = [];
        enrollAttemptsLeft = 3;
        attemptsLeftEl.textContent = '3';
        fillEl.style.width = '0%';
    }
}

function resetKeystrokeEnrollment() {
    enrolledEntries = [];
    enrollAttemptsLeft = 3;
    document.getElementById('enroll-attempts-left').textContent = '3';
    document.getElementById('enroll-progress-fill').style.width = '0%';
    document.getElementById('enroll-status').textContent = 'Training reset. Ready.';
    document.getElementById('enroll-status').style.color = 'var(--text-muted)';
    document.getElementById('enroll-keystroke-input').value = '';
    resetKeystrokeTracking();
}

// Deep Security: Steganography LSB
async function encodeStegoImage() {
    const payloadInput = document.getElementById('stego-payload');
    const fileInput = document.getElementById('stego-file-input');
    const secretText = payloadInput.value;

    if (!secretText) {
        alert('Please enter a secret payload to hide.');
        return;
    }

    const formData = new FormData();
    formData.append('secret', secretText);
    if (fileInput.files[0]) {
        formData.append('image', fileInput.files[0]);
    }

    try {
        const res = await fetch(`${API_BASE}/stego/encode`, {
            method: 'POST',
            body: formData
        });

        if (res.ok) {
            const blob = await res.blob();
            const downloadUrl = URL.createObjectURL(blob);
            
            const a = document.createElement('a');
            a.href = downloadUrl;
            a.download = 'stego_vault.png';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            
            alert('Stego PNG generated successfully! The carrier image containing your hidden secret has been downloaded.');
        } else {
            const text = await res.text();
            alert(`Stego encoding failed: ${text}`);
        }
    } catch(e) {
        alert(`Request failed: ${e.message}`);
    }
}

async function decodeStegoImage() {
    const fileInput = document.getElementById('stego-file-input');
    const resultBox = document.getElementById('stego-result-box');
    const extractedEl = document.getElementById('stego-extracted-text');

    if (!fileInput.files[0]) {
        alert('Please select an encoded stego image (PNG) to decode.');
        return;
    }

    const formData = new FormData();
    formData.append('image', fileInput.files[0]);

    try {
        const res = await fetch(`${API_BASE}/stego/decode`, {
            method: 'POST',
            body: formData
        });

        const data = await res.json();
        if (res.ok && data.status === 'success') {
            extractedEl.textContent = data.secret;
            resultBox.classList.remove('hidden');
        } else {
            alert(`Decoding failed: ${data.reason || 'Image is corrupted or contains no stego carrier payload'}`);
        }
    } catch(e) {
        alert(`Request failed: ${e.message}`);
    }
}
