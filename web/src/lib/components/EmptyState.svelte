<script lang="ts">
</script>

<div class="empty-state">
	<div class="illustration">
		<svg width="120" height="120" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
			<rect x="2" y="3" width="20" height="14" rx="2" ry="2"></rect>
			<line x1="8" y1="21" x2="16" y2="21"></line>
			<line x1="12" y1="17" x2="12" y2="21"></line>
			<line x1="8" y1="8" x2="8" y2="12"></line>
			<line x1="12" y1="8" x2="12" y2="12"></line>
			<line x1="16" y1="8" x2="16" y2="12"></line>
		</svg>
	</div>
	
	<h2>No configured service.</h2>
	<p class="subtitle">It looks like Hops hasn't been configured with any services to monitor yet.</p>

	<div class="guide">
		<h3>Quick Start Guide</h3>
		<p>Create a <code>config.yaml</code> file and point the <code>CONFIG_PATH</code> environment variable to it.</p>
		
		<div class="code-block">
			<pre><code>services:
  - name: API Gateway
    description: Main entry point
    emoji: 🚀
    checks:
      - name: health
        type: http
        target: http://api:8080/health</code></pre>
		</div>
		<p>Then restart Hops to see your services here!</p>
	</div>
</div>

<style>
	.empty-state {
		grid-column: 1 / -1;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 4rem 2rem;
		text-align: center;
		background: white;
		border-radius: 12px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
		border: 1px dashed #ced4da;
	}

	.illustration {
		color: #adb5bd;
		margin-bottom: 1.5rem;
	}

	h2 {
		margin: 0 0 0.5rem 0;
		color: #212529;
	}

	.subtitle {
		color: #6c757d;
		margin-bottom: 2rem;
		max-width: 500px;
	}

	.guide {
		text-align: left;
		background: #f8f9fa;
		padding: 1.5rem;
		border-radius: 8px;
		max-width: 600px;
		width: 100%;
	}

	h3 {
		margin: 0 0 1rem 0;
		color: #495057;
		font-size: 1.1rem;
	}

	p {
		color: #495057;
		margin-bottom: 1rem;
		line-height: 1.5;
	}

	code {
		background: #e9ecef;
		padding: 0.2rem 0.4rem;
		border-radius: 4px;
		font-family: monospace;
		font-size: 0.9em;
	}

	.code-block {
		background: #212529;
		color: #f8f9fa;
		padding: 1rem;
		border-radius: 6px;
		overflow-x: auto;
		margin-bottom: 1rem;
	}

	.code-block pre {
		margin: 0;
	}

	.code-block code {
		background: none;
		padding: 0;
		color: inherit;
	}
</style>
