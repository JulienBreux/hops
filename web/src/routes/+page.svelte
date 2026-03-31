<script lang="ts">
        import { onMount } from 'svelte';
        import { health, isEmpty } from '$lib/stores/health';
        import EmptyState from '$lib/components/EmptyState.svelte';

        onMount(() => {
                health.fetchStatus();
                const interval = setInterval(() => {
                        health.fetchStatus();
                }, 5000);

                return () => clearInterval(interval);
        });
</script>

<main class="dashboard">
        <header>
                <h1>⚡ Hops Dashboard</h1>
        </header>

        <div class="services-grid">
                {#if $isEmpty}
                        <EmptyState />
                {:else}
                        {#each $health as service}				<div class="service-card" class:down={service.checks.some((c) => !c.up)}>
					<div class="service-header">
						<span class="emoji">{service.emoji}</span>
						<h2>{service.name}</h2>
					</div>
					<p>{service.description}</p>
					<div class="checks">
						{#each service.checks as check}
							<div class="check" class:up={check.up} class:down={!check.up}>
								<span class="status-indicator"></span>
								{check.name} ({check.latency}ms)
							</div>
						{/each}
					</div>
				</div>
			{/each}
		{/if}
	</div>
</main>

<style>
	:global(body) {
		font-family:
			system-ui,
			-apple-system,
			sans-serif;
		background-color: #f8f9fa;
		margin: 0;
		color: #212529;
	}

	.dashboard {
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem;
	}

	header {
		margin-bottom: 2rem;
		text-align: center;
	}

	.services-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 1.5rem;
	}


	.service-card {
		background: white;
		border-radius: 12px;
		padding: 1.5rem;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
		border-left: 6px solid #28a745;
		transition: transform 0.2s;
	}

	.service-card.down {
		border-left-color: #dc3545;
	}

	.service-card:hover {
		transform: translateY(-4px);
	}

	.service-header {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		margin-bottom: 0.5rem;
	}

	.emoji {
		font-size: 2rem;
	}

	h2 {
		margin: 0;
		font-size: 1.25rem;
	}

	p {
		color: #6c757d;
		font-size: 0.9rem;
		margin-bottom: 1.5rem;
	}

	.checks {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.check {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-size: 0.85rem;
		font-weight: 500;
	}

	.status-indicator {
		width: 10px;
		height: 10px;
		border-radius: 50%;
	}

	.check.up .status-indicator {
		background-color: #28a745;
	}

	.check.down .status-indicator {
		background-color: #dc3545;
	}
</style>
