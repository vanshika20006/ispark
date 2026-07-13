<script lang="ts">
	import DevCredentials from '$lib/DevCredentials.svelte';
	import { goto } from '$app/navigation';
	import { API_BASE_URL } from '$lib/config';

	// Local interactive states using Svelte 5 runes
	let superAdminId = $state('');
	let password = $state('');
	let rememberMe = $state(false);
	let showPassword = $state(false);
	let submitting = $state(false);
	let errorMsg = $state('');

	const isFormValid = $derived(superAdminId.trim() !== '' && password !== '');

	// Feature list items for left column
	const features = [
		{
			title: 'Complete Platform Access',
			description: 'Manage all modules, users, admins and configurations.',
			icon: 'cog'
		},
		{
			title: 'Role & Permission Control',
			description: 'Assign and monitor access levels across the system.',
			icon: 'users'
		},
		{
			title: 'System Analytics',
			description: 'Track activities, reports and institutional insights.',
			icon: 'chart'
		},
		{
			title: 'Security Monitoring',
			description: 'Review logs and maintain platform security.',
			icon: 'shield'
		}
	];

	// Footer navigation links
	const footerLines = [
		'Skill, Personality Advancement & Refinement Cell',
		'International Institute of Professional Studies (IIPS)',
		'Devi Ahilya Vishwavidyalaya (DAVV), Indore'
	];

	const currentYear = new Date().getFullYear();

	async function handleLogin(event: SubmitEvent) {
		event.preventDefault();
		if (!isFormValid || submitting) return;

		submitting = true;
		errorMsg = '';

		try {
			const response = await fetch(`${API_BASE_URL}/api/admin/auth/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ admin_id: superAdminId, password })
			});

			const data = await response.json();

			if (!response.ok) {
				throw new Error(data.error || 'Invalid Super Admin ID or Password.');
			}

			// The login endpoint is shared with the admin portal, so an ordinary
			// admin must not be let through to the super admin portal.
			if (data.admin?.role !== 'superadmin') {
				throw new Error('This account does not have super admin access.');
			}

			localStorage.setItem('superadmin_token', data.access_token);
			await goto('/super-admin-portal/dashboard');
		} catch (err) {
			errorMsg = err instanceof Error ? err.message : 'Login failed. Please try again.';
		} finally {
			submitting = false;
		}
	}
</script>

<svelte:head>
	<title>Super Admin Control Portal | iSPARC</title>
	<meta
		name="description"
		content="Secure access for iSPARC super administrators to manage the complete platform, user roles, system settings, and institutional operations."
	/>
</svelte:head>

<div class="min-h-screen bg-[#F7F6F3] flex flex-col w-full font-sans">
	<!-- ==================== HEADER ==================== -->
	<header class="w-full bg-white border-b border-slate-200/80 sticky top-0 z-50 shadow-xs">
		<div class="max-w-6xl mx-auto flex items-center justify-between px-6 py-4">
			<!-- Brand Logo -->
			<a
				href="/"
				class="flex items-center gap-2 group focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 rounded-lg transition-all duration-200"
			>
				<div class="flex flex-col justify-center">
					<div
						class="flex items-baseline text-2xl font-bold tracking-tight text-slate-900 leading-none"
					>
						<span>i</span><span class="text-[#881B1B]">SPARC</span>
					</div>
				</div>
			</a>

			<!-- Back To Home Button -->
			<a
				href="/"
				class="inline-flex items-center gap-1.5 focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 rounded-md px-2 py-1 text-slate-650 hover:text-[#881B1B] transition-all duration-200"
			>
				<!-- Home Icon -->
				<svg
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					class="w-4 h-4"
					aria-hidden="true"
				>
					<path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
					<polyline points="9 22 9 12 15 12 15 22" />
				</svg>
				<span class="text-[13px] font-semibold uppercase tracking-wider"> Back To Home </span>
			</a>
		</div>
	</header>

	<!-- ==================== MAIN CONTENT ==================== -->
	<main class="flex-grow w-full flex items-center justify-center py-12 px-6 sm:px-8 lg:px-12">
		<div class="max-w-6xl w-full grid grid-cols-1 lg:grid-cols-12 gap-10 lg:gap-16 items-center">
			<!-- Left Column: Guidelines & Information -->
			<section class="lg:col-span-6 flex flex-col gap-6 animate-fade-in">
				<div class="flex flex-col gap-4">
					<h1
						class="text-4xl sm:text-5xl font-extrabold text-[#881B1B] font-serif leading-tight tracking-tight"
					>
						Super Admin<br />
						<span class="text-slate-900">Control Portal</span>
					</h1>
					<p class="text-slate-500 text-[14px] leading-relaxed max-w-lg">
						Secure access for iSPARC super administrators to manage the complete platform, user
						roles, system settings, and institutional operations.
					</p>
				</div>

				<!-- Feature list -->
				<div class="flex flex-col gap-6 mt-4">
					{#each features as feature}
						<div class="flex items-start gap-4 group">
							<div class="shrink-0 text-[#881B1B] mt-0.5">
								{#if feature.icon === 'cog'}
									<!-- Cog SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-5 h-5"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.43l-1.003.828c-.293.241-.438.613-.43.992a7.723 7.723 0 010 .255c-.008.378.137.75.43.991l1.004.827c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.43l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.991l-1.004-.827a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.28z"
										/>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
										/>
									</svg>
								{:else if feature.icon === 'users'}
									<!-- Users SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-5 h-5"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M18 18.72a9.094 9.094 0 003.741-.479 3 3 0 00-4.682-2.72m.94 3.198l.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0112 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 0118 18.72zm-12 0a6.062 6.062 0 0112 0v.318c0 .243-.016.481-.049.714A11.944 11.944 0 0112 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 016 18.72zm4.5-9a3 3 0 11-6 0 3 3 0 016 0zM18 9.75a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"
										/>
									</svg>
								{:else if feature.icon === 'chart'}
									<!-- Chart SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-5 h-5"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z"
										/>
									</svg>
								{:else if feature.icon === 'shield'}
									<!-- Shield SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-5 h-5"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285z"
										/>
									</svg>
								{/if}
							</div>
							<div>
								<h3
									class="text-[15px] font-bold text-slate-800 leading-snug tracking-wide group-hover:text-[#881B1B] transition-colors duration-200"
								>
									{feature.title}
								</h3>
								<p class="text-xs text-slate-500 mt-1 leading-relaxed">
									{feature.description}
								</p>
							</div>
						</div>
					{/each}
				</div>
			</section>

			<!-- Right Column: Login Card -->
			<section class="lg:col-span-6 flex justify-center lg:justify-end">
				<div
					class="w-full max-w-md bg-white rounded-2xl border border-slate-100 shadow-xl p-8 sm:p-10 transition-all duration-350 hover:shadow-2xl animate-fade-up"
				>
					<!-- Card Header -->
					<div class="pb-6 border-b border-slate-100">
						<h2 class="text-2xl font-bold text-[#0b1535] font-serif leading-tight">
							Welcome Super Admin
						</h2>
						<p class="text-slate-500 text-xs mt-1.5 leading-relaxed">
							Login to access the iSPARC master administration panel
						</p>
					</div>

					<!-- Form -->
					<form onsubmit={handleLogin} class="pt-6 flex flex-col gap-5">
						<!-- Super Admin ID Field -->
						<div class="flex flex-col gap-1.5">
							<label
								for="super-admin-id"
								class="text-[11px] font-bold text-slate-700 tracking-wider"
							>
								SUPER ADMIN ID
							</label>
							<div class="relative flex items-center">
								<span class="absolute left-3.5 pointer-events-none text-slate-400">
									<!-- Lock SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-4 h-4"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"
										/>
									</svg>
								</span>
								<input
									id="super-admin-id"
									type="text"
									bind:value={superAdminId}
									placeholder="Enter your super admin ID"
									class="w-full pl-11 pr-3.5 py-2.5 bg-white rounded-lg border border-slate-200 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
									required
								/>
							</div>
						</div>

						<!-- Password Field -->
						<div class="flex flex-col gap-1.5">
							<label for="password" class="text-[11px] font-bold text-slate-700 tracking-wider">
								PASSWORD
							</label>
							<div class="relative flex items-center">
								<span class="absolute left-3.5 pointer-events-none text-slate-400">
									<!-- Lock SVG -->
									<svg
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke-width="2.2"
										stroke="currentColor"
										class="w-4 h-4"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"
										/>
									</svg>
								</span>
								<input
									id="password"
									type={showPassword ? 'text' : 'password'}
									bind:value={password}
									placeholder="Enter your password"
									class="w-full pl-11 pr-12 py-2.5 bg-white rounded-lg border border-slate-200 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
									required
								/>
								<button
									type="button"
									onclick={() => (showPassword = !showPassword)}
									class="absolute right-3.5 text-slate-400 hover:text-slate-650 transition-colors focus:outline-none"
									aria-label={showPassword ? 'Hide password' : 'Show password'}
								>
									{#if showPassword}
										<!-- Eye-off SVG -->
										<svg
											xmlns="http://www.w3.org/2000/svg"
											fill="none"
											viewBox="0 0 24 24"
											stroke-width="2"
											stroke="currentColor"
											class="w-5 h-5"
										>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												d="M3.98 8.223A10.477 10.477 0 001.934 12C3.226 16.338 7.244 19.5 12 19.5c.993 0 1.953-.138 2.863-.395M21 21l-2.109-2.109m0 0A5.002 5.002 0 0021 12c-1.283-4.338-5.3-7.5-10.05-7.5-2.012 0-3.886.565-5.482 1.549M12 9a3 3 0 100 6M12 9a3 3 0 00-3 3M12 15a3 3 0 01-3-3"
											/>
										</svg>
									{:else}
										<!-- Eye SVG -->
										<svg
											xmlns="http://www.w3.org/2000/svg"
											fill="none"
											viewBox="0 0 24 24"
											stroke-width="2"
											stroke="currentColor"
											class="w-5 h-5"
										>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.43 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z"
											/>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
											/>
										</svg>
									{/if}
								</button>
							</div>
						</div>

						<!-- Remember Me / Forgot Password -->
						<div class="flex items-center justify-between mt-1">
							<label class="flex items-center gap-2 cursor-pointer select-none">
								<input
									type="checkbox"
									bind:checked={rememberMe}
									class="w-4 h-4 rounded border-slate-355 text-[#881B1B] focus:ring-[#881B1B] focus:ring-opacity-20 cursor-pointer accent-[#881B1B]"
								/>
								<span
									class="text-xs text-slate-650 font-medium hover:text-slate-800 transition-colors"
								>
									Remember Me
								</span>
							</label>
							<a
								href="#forgot"
								onclick={(e) => {
									e.preventDefault();
									alert(
										'Please contact the system administrator to retrieve or reset your super admin credentials.'
									);
								}}
								class="text-xs font-semibold text-[#881B1B] hover:underline transition-colors"
							>
								Forgot Password?
							</a>
						</div>

						<!-- Error message -->
						{#if errorMsg}
							<div
								class="p-3 bg-red-50 border border-red-200 rounded-lg text-xs font-semibold text-red-700"
								role="alert"
							>
								{errorMsg}
							</div>
						{/if}

						<!-- Submit Button -->
						<button
							type="submit"
							disabled={!isFormValid || submitting}
							class="w-full py-3 bg-[#0B1535] hover:bg-[#132049] text-white font-bold text-xs tracking-widest uppercase rounded-lg transition-colors duration-200 shadow-sm mt-3 disabled:opacity-50 disabled:cursor-not-allowed"
						>
							{submitting ? 'SIGNING IN…' : 'LOGIN'}
						</button>

						<!-- Security Notice Box -->
						<div
							class="flex items-start gap-3.5 p-4 bg-red-50 border border-red-100 rounded-xl mt-2"
						>
							<div class="shrink-0 mt-0.5 text-red-700">
								<!-- Shield Warning SVG -->
								<svg
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
									stroke-width="2"
									stroke="currentColor"
									class="w-4 h-4"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
									/>
								</svg>
							</div>
							<div>
								<h4 class="text-xs font-bold text-red-800 leading-tight">Security Notice</h4>
								<p class="text-[11px] text-red-700 mt-1 leading-normal font-medium font-sans">
									Only authorized super administrators can access this portal.
								</p>
							</div>
						</div>
					</form>
					<DevCredentials portal="superadmin" />
				</div>
			</section>
		</div>
	</main>

	<!-- ==================== FOOTER ==================== -->
	<footer class="w-full bg-[#0b1535] py-8 border-t border-slate-900 mt-auto">
		<div class="max-w-6xl mx-auto flex flex-col items-center gap-1.5 px-6 text-center">
			<h2 class="text-base font-bold text-white tracking-wider uppercase font-sans">iSPARC</h2>
			{#each footerLines as line}
				<p class="text-xs font-normal text-white/50 leading-relaxed font-sans">
					{line}
				</p>
			{/each}
			<p class="text-[10px] text-white/30 mt-2 font-sans">
				© {currentYear} IIPS DAVV. All rights reserved.
			</p>
		</div>
	</footer>
</div>
