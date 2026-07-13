<script lang="ts">
	import { slide } from 'svelte/transition';

	let { portal }: { portal: 'student' | 'admin' | 'superadmin' } = $props();

	const DEV_PASSWORD = 'Pass@123';

	const accounts: Record<string, { label: string; login: string; note: string }[]> = {
		student: [
			{ label: 'Rahul Sharma', login: 'rahul.sharma@iips.edu', note: 'IT2K24 · most data' },
			{ label: 'Sneha Kumar', login: 'sneha.kumar@iips.edu', note: 'IT2K24' },
			{ label: 'Arjun Desai', login: 'arjun.desai@iips.edu', note: 'IT2K24 · has a rejection' },
			{ label: 'Priya Nair', login: 'priya.nair@iips.edu', note: 'IT2K25' },
			{ label: 'Vikram Singh', login: 'vikram.singh@iips.edu', note: 'IT2K24 · no activity' }
		],
		admin: [
			{ label: 'Dr. Rajesh Kumar', login: 'admin', note: 'batch IT2K24' },
			{ label: 'Dr. Priya Patel', login: 'admin2', note: 'batch IT2K25' }
		],
		superadmin: [{ label: 'Ananya Deshmukh', login: 'superadmin', note: 'whole platform' }]
	};

	let open = $state(false);
	let copied = $state('');

	const list = $derived(accounts[portal] ?? []);

	async function copy(value: string) {
		try {
			await navigator.clipboard.writeText(value);
			copied = value;
			setTimeout(() => (copied = ''), 1200);
		} catch {
			copied = '';
		}
	}
</script>

<div class="mt-5 border border-amber-200 bg-amber-50/60 rounded-xl overflow-hidden">
	<button
		type="button"
		onclick={() => (open = !open)}
		class="w-full flex items-center justify-between gap-3 px-4 py-2.5 text-left"
	>
		<span class="flex items-center gap-2">
			<span
				class="px-1.5 py-0.5 rounded bg-amber-500 text-white text-[9px] font-extrabold tracking-widest uppercase"
			>
				Dev
			</span>
			<span class="text-[11px] font-bold text-amber-900">Demo login credentials</span>
		</span>
		<span class="text-[10px] font-bold text-amber-700 uppercase tracking-wider">
			{open ? 'Hide' : 'Show'}
		</span>
	</button>

	{#if open}
		<div transition:slide={{ duration: 150 }} class="px-4 pb-4 flex flex-col gap-2">
			<p class="text-[10px] text-amber-800 leading-relaxed">
				Seeded demo accounts. Every one of them uses the password
				<button
					type="button"
					onclick={() => copy(DEV_PASSWORD)}
					class="font-mono font-bold underline decoration-dotted hover:text-amber-950"
				>
					{DEV_PASSWORD}
				</button>. Click any value to copy it.
			</p>

			<div class="flex flex-col gap-1.5">
				{#each list as account (account.login)}
					<div class="flex items-center justify-between gap-3 bg-white/70 rounded-lg px-3 py-2">
						<div class="min-w-0">
							<button
								type="button"
								onclick={() => copy(account.login)}
								class="block text-[11px] font-mono font-bold text-slate-800 truncate hover:underline"
							>
								{account.login}
							</button>
							<span class="text-[9px] text-slate-500">{account.label} · {account.note}</span>
						</div>
						{#if copied === account.login}
							<span class="text-[9px] font-bold text-emerald-600 shrink-0">Copied</span>
						{/if}
					</div>
				{/each}
			</div>

			{#if copied === DEV_PASSWORD}
				<span class="text-[9px] font-bold text-emerald-600">Password copied</span>
			{/if}
		</div>
	{/if}
</div>
