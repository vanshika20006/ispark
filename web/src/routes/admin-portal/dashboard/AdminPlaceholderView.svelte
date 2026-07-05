<script lang="ts">
	import { fade } from 'svelte/transition';

	// Props using Svelte 5 runes
	let {
		tabName = 'Feature',
		onBackToDashboard
	}: {
		tabName: string;
		onBackToDashboard: () => void;
	} = $props();

	// Map tab name to description
	const tabDetails = $derived(() => {
		switch (tabName) {
			case 'Student Management':
				return {
					description:
						'Browse all registered students, manage their profiles, tracks, and track total credit allocations.'
				};
			case 'Certificate Verification':
				return {
					description:
						'Verify, reject, and review certificate uploads submitted by students across all branches.'
				};
			case 'Activity Monitoring':
				return {
					description:
						'Track and monitor real-time activity participation logs, hourly contributions, and enrollment statuses.'
				};
			case 'Batch Analytics':
				return {
					description:
						'View aggregated performance charts, track credit completion ratios, and generate cohort analytics.'
				};
			default:
				return {
					description:
						'This feature is currently under development as part of the Phase 3 Admin Module.'
				};
		}
	});
</script>

<div
	transition:fade={{ duration: 150 }}
	class="bg-white border border-border-base rounded-xl p-8 sm:p-12 text-center max-w-2xl mx-auto my-12 shadow-xs space-y-6"
>
	<div
		class="w-20 h-20 bg-slate-50 border border-slate-100 flex items-center justify-center rounded-full mx-auto"
	>
		{#if tabName === 'Student Management'}
			<svg
				class="w-16 h-16 text-slate-400 animate-pulse"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="1.5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M15 19.128a9.38 9.38 0 0 0 2.625.372 9.337 9.337 0 0 0 4.121-.952 4.125 4.125 0 0 0-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.109A11.386 11.386 0 0 1 10.089 21c-2.316 0-4.445-.69-6.22-1.879v-.003a4.125 4.125 0 0 1 7.533-2.493M15 19.128v-.003c0-1.112-.285-2.16-.786-3.07M14.214 16.058A9.396 9.396 0 0 0 10.089 15c-1.47 0-2.854.34-4.082.945M14.214 16.058a9.386 9.386 0 0 1 0 3.07M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
				/>
			</svg>
		{:else if tabName === 'Certificate Verification'}
			<svg
				class="w-16 h-16 text-slate-400 animate-pulse"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="1.5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M9 12.75 11.25 15 15 9.75M21 12c0 1.268-.63 2.39-1.593 3.068a3.745 3.745 0 0 1-1.043 3.296 3.745 3.745 0 0 1-3.296 1.043A3.745 3.745 0 0 1 12 21c-1.268 0-2.39-.63-3.068-1.593a3.746 3.746 0 0 1-3.296-1.043 3.745 3.745 0 0 1-1.043-3.296A3.745 3.745 0 0 1 3 12c0-1.268.63-2.39 1.593-3.068a3.745 3.745 0 0 1 1.043-3.296 3.746 3.746 0 0 1 3.296-1.043A3.746 3.746 0 0 1 12 3c1.268 0 2.39.63 3.068 1.593a3.746 3.746 0 0 1 3.296 1.043 3.746 3.746 0 0 1 1.043 3.296A3.745 3.745 0 0 1 21 12Z"
				/>
			</svg>
		{:else if tabName === 'Activity Monitoring'}
			<svg
				class="w-16 h-16 text-slate-400 animate-pulse"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="1.5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M2.25 18 9 11.25l4.306 4.306a11.95 11.95 0 0 1 5.814-5.518l2.74-1.22m0 0-5.94-2.281m5.94 2.28-2.28 5.941"
				/>
			</svg>
		{:else if tabName === 'Batch Analytics'}
			<svg
				class="w-16 h-16 text-slate-400 animate-pulse"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="1.5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M10.5 6a7.5 7.5 0 1 0 7.5 7.5h-7.5V6Z"
				/>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M13.5 10.5H21A7.5 7.5 0 0 0 13.5 3v7.5Z"
				/>
			</svg>
		{:else}
			<svg
				class="w-16 h-16 text-slate-400 animate-pulse"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				stroke-width="1.5"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
				/>
			</svg>
		{/if}
	</div>

	<div class="space-y-2">
		<h2 class="text-xl font-bold font-serif text-inst-navy">{tabName}</h2>
		<p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest">
			Phase 3 &middot; Under Development
		</p>
		<p class="text-sm text-slate-500 max-w-md mx-auto leading-relaxed pt-2">
			{tabDetails().description}
		</p>
	</div>

	<div class="pt-4 flex justify-center gap-4">
		<button
			onclick={onBackToDashboard}
			class="px-5 py-2.5 bg-accent-red hover:bg-accent-red/90 text-white font-bold text-xs tracking-wider uppercase rounded-lg transition duration-200 shadow-xs focus:outline-none focus:ring-2 focus:ring-accent-red/20"
		>
			Back to Dashboard
		</button>
	</div>
</div>
