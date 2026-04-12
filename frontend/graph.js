// A custom, hand-written Force-Directed Graph Engine using HTML5 Canvas

class NeuralNexus {
    constructor(canvasId) {
        this.canvas = document.getElementById(canvasId);
        this.ctx = this.canvas.getContext('2d');
        this.nodes = [];
        this.edges = [];
        
        // Physics constants
        this.repulsion = 2000;
        this.springLength = 80;
        this.springK = 0.05;
        this.damping = 0.85;
        
        this.width = this.canvas.width;
        this.height = this.canvas.height;
        
        this.animationId = null;

        this.resize();
        window.addEventListener('resize', () => this.resize());
    }

    resize() {
        const parent = this.canvas.parentElement;
        if (parent) {
            this.width = parent.clientWidth;
            this.height = parent.clientHeight;
            this.canvas.width = this.width;
            this.canvas.height = this.height;
        }
    }

    loadData(graphData) {
        // Map nodes
        const nodeMap = new Map();
        this.nodes = graphData.nodes.map(n => {
            const node = {
                id: n.id,
                label: n.label,
                group: n.group,
                size: n.size,
                x: Math.random() * this.width,
                y: Math.random() * this.height,
                vx: 0,
                vy: 0
            };
            nodeMap.set(n.id, node);
            return node;
        });

        // Map edges
        this.edges = graphData.edges.map(e => ({
            source: nodeMap.get(e.source),
            target: nodeMap.get(e.target)
        })).filter(e => e.source && e.target);
    }

    applyPhysics() {
        // Coulomb's Law (Repulsion)
        for (let i = 0; i < this.nodes.length; i++) {
            for (let j = i + 1; j < this.nodes.length; j++) {
                const n1 = this.nodes[i];
                const n2 = this.nodes[j];
                
                const dx = n1.x - n2.x;
                const dy = n1.y - n2.y;
                let dist = Math.sqrt(dx*dx + dy*dy);
                if (dist === 0) dist = 0.1;

                if (dist < 300) { // Limit repulsion range for performance
                    const force = this.repulsion / (dist * dist);
                    const fx = (dx / dist) * force;
                    const fy = (dy / dist) * force;

                    n1.vx += fx;
                    n1.vy += fy;
                    n2.vx -= fx;
                    n2.vy -= fy;
                }
            }
        }

        // Hooke's Law (Spring Attraction)
        for (const edge of this.edges) {
            const dx = edge.target.x - edge.source.x;
            const dy = edge.target.y - edge.source.y;
            let dist = Math.sqrt(dx*dx + dy*dy);
            if (dist === 0) dist = 0.1;

            const displacement = dist - this.springLength;
            const force = this.springK * displacement;
            
            const fx = (dx / dist) * force;
            const fy = (dy / dist) * force;

            edge.source.vx += fx;
            edge.source.vy += fy;
            edge.target.vx -= fx;
            edge.target.vy -= fy;
        }

        // Center Gravity (Keep things on screen)
        const cx = this.width / 2;
        const cy = this.height / 2;
        const gravity = 0.02;

        for (const node of this.nodes) {
            node.vx += (cx - node.x) * gravity;
            node.vy += (cy - node.y) * gravity;

            // Integration & Damping
            node.vx *= this.damping;
            node.vy *= this.damping;
            
            node.x += node.vx;
            node.y += node.vy;
        }
    }

    render() {
        this.ctx.clearRect(0, 0, this.width, this.height);

        // Draw edges
        this.ctx.lineWidth = 1;
        this.ctx.strokeStyle = "rgba(180, 180, 180, 0.3)";
        for (const edge of this.edges) {
            this.ctx.beginPath();
            this.ctx.moveTo(edge.source.x, edge.source.y);
            this.ctx.lineTo(edge.target.x, edge.target.y);
            this.ctx.stroke();
        }

        // Draw nodes
        for (const node of this.nodes) {
            this.ctx.beginPath();
            this.ctx.arc(node.x, node.y, node.size, 0, 2 * Math.PI);
            
            if (node.group === 1) { // Document
                this.ctx.fillStyle = "#0071e3"; // Blue
            } else { // Tag
                this.ctx.fillStyle = "#34c759"; // Green
            }
            
            this.ctx.fill();
            this.ctx.strokeStyle = "#fff";
            this.ctx.lineWidth = 2;
            this.ctx.stroke();

            // Label
            if (node.group === 2 || node.size > 10) {
                this.ctx.fillStyle = "#333";
                this.ctx.font = node.group === 2 ? "12px Inter" : "10px Inter";
                this.ctx.fillText(node.label, node.x + node.size + 4, node.y + 4);
            }
        }
    }

    loop() {
        this.applyPhysics();
        this.render();
        this.animationId = requestAnimationFrame(() => this.loop());
    }

    start() {
        if (!this.animationId) {
            this.loop();
        }
    }

    stop() {
        if (this.animationId) {
            cancelAnimationFrame(this.animationId);
            this.animationId = null;
        }
    }
}
