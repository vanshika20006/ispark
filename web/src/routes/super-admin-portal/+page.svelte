<script lang="ts">
	import { fade, slide } from 'svelte/transition';

	// Sidebar menu items list for Super Admin Portal
	const menuItems = [
		{
			name: 'Dashboard',
			icon: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6'
		},
		{
			name: 'User Management',
			icon: 'M15 19.128a9.38 9.38 0 0 0 2.625.372 9.337 9.337 0 0 0 4.121-.952 4.125 4.125 0 0 0-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.109A11.386 11.386 0 0 1 10.089 21c-2.316 0-4.445-.69-6.22-1.879v-.003a4.125 4.125 0 0 1 7.533-2.493M15 19.128v-.003c0-1.112-.285-2.16-.786-3.07M14.214 16.058A9.396 9.396 0 0 0 10.089 15c-1.47 0-2.854.34-4.082.945M14.214 16.058a9.386 9.386 0 0 1 0 3.07M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5'
		},
		{
			name: 'Activity Management',
			icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z'
		},
		{
			name: 'Track Management',
			icon: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z'
		},
		{
			name: 'Announcement Management',
			icon: 'M19.114 5.636a9 9 0 010 12.728M16.463 8.288a5.25 5.25 0 010 7.424M6.75 8.25l4.72-4.72a.75.75 0 011.28.53v15.88a.75.75 0 01-1.28.53l-4.72-4.72H4.51c-.88 0-1.704-.507-1.938-1.354A9.01 9.01 0 012.25 12c0-.83.112-1.633.322-2.396C2.806 8.756 3.63 8.25 4.51 8.25H6.75z'
		},
		{
			name: 'Reports Center',
			icon: 'M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z'
		},
		{
			name: 'System Settings',
			icon: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z'
		}
	];

	// Selected menu state
	let currentTab = $state('Dashboard');
	let isMobileSidebarOpen = $state(false);

	// Search & Notification states
	let isNotificationsOpen = $state(false);
	let searchQuery = $state('');

	// Mock Notifications for Super Admin
	const notifications = [
		{
			id: 1,
			text: 'System health warning: Background DB sync delayed by 3m.',
			time: '15 mins ago',
			unread: true
		},
		{
			id: 2,
			text: 'Security Alert: Failed SSH attempt from 192.168.1.104.',
			time: '1 hour ago',
			unread: true
		},
		{
			id: 3,
			text: 'Daily audit backup successfully uploaded to S3.',
			time: '6 hours ago',
			unread: false
		}
	];

	// Mock Recent Users data (Step 3)
	const recentUsers = [
		{ name: 'Rahul Sharma', id: 'ENR024001', role: 'Student', dept: 'Computer Science', status: 'Active' },
		{ name: 'Dr. Priya Patel', id: 'ADM024010', role: 'Admin', dept: 'Electronics & Comm.', status: 'Active' },
		{ name: 'Arjun Desai', id: 'ENR024008', role: 'Student', dept: 'Data Science', status: 'Pending' },
		{ name: 'Sneha Kumar', id: 'ENR024015', role: 'Student', dept: 'Mechanical Engg.', status: 'Active' },
		{ name: 'Dr. Vikram Singh', id: 'ADM024003', role: 'Admin', dept: 'Civil Engineering', status: 'Active' }
	];

	function toggleMobileSidebar() {
		isMobileSidebarOpen = !isMobileSidebarOpen;
	}

	function handleLogout() {
		// Super Admin Logout -> redirect to home page
		window.location.href = '/';
	}
</script>

<svelte:head>
	<title>Super Admin Console | iSPARC</title>
	<meta
		name="description"
		content="iSPARC Super Admin Portal. Configure tracks, view activity logs, manage administration settings, and oversee platform status."
	/>
</svelte:head>

<div class="min-h-screen bg-[#F7F6F3] text-slate-800 flex font-sans">
	<!-- ==================== SIDEBAR ==================== -->
	<!-- Desktop Sidebar -->
	<aside class="hidden lg:flex flex-col w-64 bg-white border-r border-slate-200 h-screen sticky top-0 shrink-0">
		<!-- Logo area -->
		<div class="h-[72px] flex items-center px-6 gap-3.5 border-b border-slate-100">
			<!-- Red square badge with white SA -->
			<div class="w-9 h-9 bg-[#881B1B] rounded-lg flex items-center justify-center text-white font-bold font-serif text-[11px] shrink-0">
				SA
			</div>
			<div class="flex flex-col">
				<span class="text-xl font-bold tracking-tight text-slate-900 font-serif">
					i<span class="text-[#881B1B]">SPARC</span>
				</span>
			</div>
		</div>

		<!-- Menu Items -->
		<nav class="flex-grow py-6 px-4 space-y-1 overflow-y-auto">
			{#each menuItems as item}
				<button
					onclick={() => (currentTab = item.name)}
					class="w-full flex items-center gap-3.5 px-4 py-3 rounded-lg text-[13px] font-bold tracking-wide transition-all duration-200 {currentTab === item.name
						? 'bg-[#881B1B]/10 text-[#881B1B] border-l-[3px] border-[#881B1B] rounded-l-none'
						: 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'}"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
						class="w-5 h-5"
					>
						<path stroke-linecap="round" stroke-linejoin="round" d={item.icon} />
					</svg>
					{item.name}
				</button>
			{/each}
		</nav>

		<!-- Logout -->
		<div class="p-4 border-t border-slate-100">
			<button
				onclick={handleLogout}
				class="w-full flex items-center gap-3 px-4 py-3 rounded-lg text-[13px] font-bold text-slate-500 hover:bg-rose-50 hover:text-rose-600 transition-colors"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="2"
					class="w-5 h-5"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
					/>
				</svg>
				Logout
			</button>
		</div>
	</aside>

	<!-- Mobile Sidebar Overlay Menu (Drawer) -->
	{#if isMobileSidebarOpen}
		<!-- Backdrop -->
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			onclick={toggleMobileSidebar}
			transition:fade={{ duration: 150 }}
			class="lg:hidden fixed inset-0 bg-black/40 z-40"
		></div>

		<!-- Drawer Content -->
		<aside
			transition:slide={{ axis: 'x', duration: 250 }}
			class="lg:hidden fixed left-0 top-0 bottom-0 w-64 bg-white z-50 flex flex-col h-full shadow-2xl"
		>
			<div class="h-[72px] flex items-center justify-between px-6 border-b border-slate-100">
				<div class="flex items-center gap-3">
					<div class="w-9 h-9 bg-[#881B1B] rounded-lg flex items-center justify-center text-white font-bold font-serif text-[11px] shrink-0">
						SA
					</div>
					<div class="flex flex-col">
						<span class="text-xl font-bold tracking-tight text-slate-900 font-serif">
							i<span class="text-[#881B1B]">SPARC</span>
						</span>
					</div>
				</div>
				<button
					onclick={toggleMobileSidebar}
					aria-label="Close sidebar"
					class="p-1 rounded-lg text-slate-400 hover:bg-slate-100 hover:text-slate-600"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
						class="w-6 h-6"
					>
						<path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
					</svg>
				</button>
			</div>

			<nav class="flex-grow py-6 px-4 space-y-1 overflow-y-auto">
				{#each menuItems as item}
					<button
						onclick={() => {
							currentTab = item.name;
							toggleMobileSidebar();
						}}
						class="w-full flex items-center gap-3.5 px-4 py-3 rounded-lg text-[13px] font-bold tracking-wide transition-all duration-200 {currentTab === item.name
							? 'bg-[#881B1B]/10 text-[#881B1B] border-l-[3px] border-[#881B1B] rounded-l-none'
							: 'text-slate-500 hover:bg-slate-50 hover:text-slate-900'}"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="2"
							class="w-5 h-5"
						>
							<path stroke-linecap="round" stroke-linejoin="round" d={item.icon} />
						</svg>
						{item.name}
					</button>
				{/each}
			</nav>

			<div class="p-4 border-t border-slate-100">
				<button
					onclick={() => {
						handleLogout();
						toggleMobileSidebar();
					}}
					class="w-full flex items-center gap-3 px-4 py-3 rounded-lg text-[13px] font-bold text-slate-500 hover:bg-rose-50 hover:text-rose-600 transition-colors"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
						class="w-5 h-5"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
						/>
					</svg>
					Logout
				</button>
			</div>
		</aside>
	{/if}

	<!-- ==================== MAIN WORKSPACE ==================== -->
	<div class="flex-grow flex flex-col min-w-0">
		<!-- Top Navigation Header -->
		<header class="bg-white border-b border-slate-200 h-[72px] flex items-center justify-between px-4 sm:px-6 lg:px-8 sticky top-0 z-30">
			<!-- Mobile sidebar toggle -->
			<div class="flex items-center gap-3">
				<button
					onclick={toggleMobileSidebar}
					aria-label="Open sidebar"
					class="lg:hidden p-2 rounded-lg text-slate-500 hover:bg-slate-100 hover:text-slate-900 focus:outline-none"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
						class="w-6 h-6"
					>
						<path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
					</svg>
				</button>
				<div>
					<h1 class="text-lg sm:text-xl font-bold text-slate-900 font-serif leading-tight">
						{currentTab === 'Dashboard' ? 'Super Admin Console' : currentTab}
					</h1>
					<p class="text-[10px] sm:text-xs font-semibold text-slate-400 uppercase tracking-wider mt-0.5">
						Tuesday, 23 June 2026
					</p>
				</div>
			</div>

			<!-- Action items -->
			<div class="flex items-center gap-3.5 relative">
				<!-- Search box -->
				<div class="hidden md:flex items-center relative">
					<input
						type="text"
						bind:value={searchQuery}
						placeholder="Search..."
						class="pl-4 pr-9 py-2 bg-slate-50 rounded-lg border border-slate-200 text-xs text-slate-800 focus:outline-none focus:border-slate-350 focus:bg-white w-52 transition-all focus:w-60"
					/>
					<span class="absolute right-3 text-slate-400">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="2"
							class="w-4 h-4"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
							/>
						</svg>
					</span>
				</div>

				<!-- Notification button -->
				<button
					onclick={() => {
						isNotificationsOpen = !isNotificationsOpen;
					}}
					class="w-10 h-10 border border-slate-250 bg-white rounded-lg flex items-center justify-center text-slate-500 hover:bg-slate-50 transition-colors relative"
				>
					<span class="sr-only">Notifications</span>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
						class="w-5 h-5"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
						/>
					</svg>
					<span class="absolute top-2 right-2.5 w-2 h-2 bg-[#881B1B] rounded-full"></span>
				</button>

				{#if isNotificationsOpen}
					<div
						transition:slide={{ duration: 150 }}
						class="absolute right-0 top-12 w-80 bg-white border border-slate-200 rounded-xl shadow-lg py-2.5 z-40 mt-1 animate-fade-in"
					>
						<h3 class="px-4 py-1.5 text-xs font-bold text-slate-400 uppercase tracking-wider border-b border-slate-100">
							System Alerts
						</h3>
						<div class="max-h-64 overflow-y-auto">
							{#each notifications as notice}
								<div class="px-4 py-3 hover:bg-slate-50 border-b border-slate-50 last:border-b-0 flex gap-2.5">
									{#if notice.unread}
										<span class="w-2.5 h-2.5 bg-[#881B1B] rounded-full mt-1.5 shrink-0"></span>
									{/if}
									<div class="flex-grow">
										<p class="text-xs text-slate-700 font-semibold">{notice.text}</p>
										<span class="text-[10px] text-slate-405 block mt-1">{notice.time}</span>
									</div>
								</div>
							{/each}
						</div>
					</div>
				{/if}

				<!-- Profile Pill -->
				<div class="flex items-center gap-3 pl-2 border-l border-slate-200 select-none">
					<div class="w-10 h-10 rounded-full bg-[#881B1B] text-white flex items-center justify-center font-bold text-sm border-2 border-white shadow-sm shrink-0 font-serif">
						SA
					</div>
					<div class="hidden sm:flex flex-col">
						<span class="text-xs font-bold text-slate-900 leading-none">Super Admin</span>
						<span class="text-[9px] text-slate-400 font-bold tracking-wider block mt-1 uppercase">
							Platform Administrator
						</span>
					</div>
				</div>
			</div>
		</header>

		<!-- Main Content Body -->
		<main class="flex-grow p-4 sm:p-6 lg:p-8 space-y-6 overflow-y-auto max-w-7xl mx-auto w-full">
			{#if currentTab === 'Dashboard'}
				<!-- Welcome Banner Section -->
				<div class="bg-white border border-slate-200 rounded-xl p-6 sm:p-8 shadow-xs space-y-2">
					<h2 class="text-2xl font-bold font-serif text-[#0B1535]">Welcome Back, Super Admin!</h2>
					<p class="text-xs text-slate-500 font-medium">
						Academic Year 2025–26 &middot; Tuesday, 23 June 2026
					</p>
				</div>
				
				<!-- Dashboard Statistics Cards Grid (Step 2) -->
				<section class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6" aria-label="Super Admin metrics overview">
					<!-- Card 1: Total Students -->
					<div class="bg-white border border-slate-200 rounded-xl p-6 shadow-xs flex flex-col justify-between hover:shadow-md transition-shadow">
						<div class="flex items-center justify-between">
							<span class="text-2xl font-bold font-serif text-slate-900">1,248</span>
							<div class="p-2.5 rounded-lg bg-blue-50 text-blue-600 border border-blue-100">
								<!-- People icon -->
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5">
									<path stroke-linecap="round" stroke-linejoin="round" d="M15 19.128a9.38 9.38 0 002.625.372 9.337 9.337 0 004.121-.952 4.125 4.125 0 00-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.109A11.386 11.386 0 0 1 10.089 21c-2.316 0-4.445-.69-6.22-1.879v-.003a4.125 4.125 0 0 1 7.533-2.493M15 19.128v-.003c0-1.112-.285-2.16-.786-3.07M14.214 16.058A9.396 9.396 0 0 0 10.089 15c-1.47 0-2.854.34-4.082.945M14.214 16.058a9.386 9.386 0 0 1 0 3.07" />
								</svg>
							</div>
						</div>
						<div class="mt-4">
							<h3 class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">Total Students</h3>
							<span class="text-[11px] font-bold text-slate-400 mt-1 block">+32 this semester</span>
						</div>
					</div>

					<!-- Card 2: Total Admins -->
					<div class="bg-white border border-slate-200 rounded-xl p-6 shadow-xs flex flex-col justify-between hover:shadow-md transition-shadow">
						<div class="flex items-center justify-between">
							<span class="text-2xl font-bold font-serif text-slate-900">87</span>
							<div class="p-2.5 rounded-lg bg-blue-50 text-blue-600 border border-blue-100">
								<!-- Single user admin icon -->
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5">
									<path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 11-7.5 0 3.75 3.75 0 017.5 0zM4.501 20.118a7.5 7.5 0 0114.998 0A17.933 17.933 0 0112 21.75c-2.676 0-5.216-.584-7.499-1.632z" />
								</svg>
							</div>
						</div>
						<div class="mt-4">
							<h3 class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">Total Admins</h3>
							<span class="text-[11px] font-bold text-slate-400 mt-1 block">+5 this month</span>
						</div>
					</div>

					<!-- Card 3: Total Activities -->
					<div class="bg-white border border-slate-200 rounded-xl p-6 shadow-xs flex flex-col justify-between hover:shadow-md transition-shadow">
						<div class="flex items-center justify-between">
							<span class="text-2xl font-bold font-serif text-slate-900">3,412</span>
							<div class="p-2.5 rounded-lg bg-blue-50 text-blue-600 border border-blue-100">
								<!-- Graph icon -->
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5">
									<path stroke-linecap="round" stroke-linejoin="round" d="M2.25 18 9 11.25l4.306 4.306a11.95 11.95 0 0 1 5.814-5.518l2.74-1.22m0 0-5.94-2.281m5.94 2.28-2.28 5.941" />
								</svg>
							</div>
						</div>
						<div class="mt-4">
							<h3 class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">Total Activities</h3>
							<span class="text-[11px] font-bold text-slate-400 mt-1 block">+110 this week</span>
						</div>
					</div>

					<!-- Card 4: Active Tracks -->
					<div class="bg-white border border-slate-200 rounded-xl p-6 shadow-xs flex flex-col justify-between hover:shadow-md transition-shadow">
						<div class="flex items-center justify-between">
							<span class="text-2xl font-bold font-serif text-slate-900">6</span>
							<div class="p-2.5 rounded-lg bg-blue-50 text-blue-600 border border-blue-100">
								<!-- Stack icon -->
								<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-5 h-5">
									<path stroke-linecap="round" stroke-linejoin="round" d="M6.429 9.75L2.25 12l4.179 2.25m0-4.5l5.571 3 5.571-3m-11.142 0L21.75 12l-4.179 2.25m0 0L21.75 16.5l-9.75 5.25-9.75-5.25L6.429 14.25m11.142 0L12 16.5m0-13.5L21.75 7.5 12 12.75 2.25 7.5 12 3z" />
								</svg>
							</div>
						</div>
						<div class="mt-4">
							<h3 class="text-[10px] font-bold text-slate-400 uppercase tracking-wider">Active Tracks</h3>
							<span class="text-[11px] font-bold text-slate-400 mt-1 block">2 new this year</span>
						</div>
					</div>
				</section>

				<!-- Middle Grid section (Step 3) -->
				<div class="grid grid-cols-1 lg:grid-cols-3 gap-6 items-start">
					<!-- User Management Overview (lg:col-span-2) -->
					<div class="lg:col-span-2 bg-white border border-slate-200 rounded-xl shadow-xs overflow-hidden">
						<div class="p-5 border-b border-slate-100 flex items-center justify-between bg-slate-50/20">
							<h3 class="text-sm font-bold font-serif text-slate-905">User Management Overview</h3>
							<button onclick={() => (currentTab = 'User Management')} class="text-[#881B1B] hover:underline text-xs font-bold uppercase tracking-wider">
								View All
							</button>
						</div>
						<div class="overflow-x-auto">
							<table class="w-full text-left border-collapse">
								<thead>
									<tr class="border-b border-slate-150 bg-slate-50/50 text-[10px] font-extrabold text-slate-405 uppercase tracking-wider">
										<th class="py-3.5 px-5">User Name</th>
										<th class="py-3.5 px-5">Role</th>
										<th class="py-3.5 px-5">Department</th>
										<th class="py-3.5 px-5">Status</th>
									</tr>
								</thead>
								<tbody class="divide-y divide-slate-100 text-xs font-sans">
									{#each recentUsers as user}
										<tr class="hover:bg-slate-50/30 transition-colors">
											<td class="py-4 px-5">
												<div class="flex flex-col">
													<span class="font-bold text-slate-800">{user.name}</span>
													<span class="text-[10px] text-slate-400 font-semibold mt-0.5 select-all">{user.id}</span>
												</div>
											</td>
											<td class="py-4 px-5">
												<span class="inline-flex items-center px-2 py-0.5 text-[10px] font-bold uppercase rounded-full border
													{user.role === 'Student'
														? 'bg-slate-50 text-slate-650 border-slate-200'
														: 'bg-blue-50 text-blue-700 border-blue-100'}"
												>
													{user.role}
												</span>
											</td>
											<td class="py-4 px-5 text-slate-500 font-semibold">{user.dept}</td>
											<td class="py-4 px-5">
												<span class="inline-flex items-center gap-1.5 font-bold
													{user.status === 'Active' ? 'text-emerald-600' : 'text-amber-600'}"
												>
													<span class="w-1.5 h-1.5 rounded-full shrink-0
														{user.status === 'Active' ? 'bg-emerald-600' : 'bg-amber-500 animate-pulse'}"
													></span>
													{user.status}
												</span>
											</td>
										</tr>
									{/each}
								</tbody>
							</table>
						</div>
					</div>

					<!-- Right Column placeholder: Quick Actions (Step 4 placeholder) -->
					<div class="bg-white border border-slate-200 rounded-xl p-5 shadow-xs space-y-4">
						<h3 class="text-sm font-bold font-serif text-slate-905">Quick Actions</h3>
						<div class="h-px bg-slate-100 my-2"></div>
						<div class="p-6 border border-dashed border-slate-350 bg-slate-50/50 rounded-xl text-center text-xs text-slate-400 leading-relaxed font-semibold">
							Quick action buttons (Create User, Create Activity, Create Track, Generate Report) will be added here in Step 4.
						</div>
					</div>
				</div>

				<!-- Recent System Activities Section Placeholder (Step 5 placeholder) -->
				<div class="p-8 border border-dashed border-slate-350 bg-white rounded-xl text-center space-y-2">
					<h3 class="font-bold text-slate-800 text-sm">System Logs & Activities Pending</h3>
					<p class="text-xs text-slate-400 max-w-sm mx-auto leading-relaxed">
						Super admin event logs detailing system activity logs (audit trailing, announcement tracking, and report dispatches) will be added here in Step 5.
					</p>
				</div>
			{:else}
				<!-- Under Construction placeholder for other tabs -->
				<div class="bg-white border border-slate-200 rounded-xl p-8 sm:p-12 text-center max-w-2xl mx-auto my-12 shadow-xs space-y-6">
					<div class="w-16 h-16 bg-[#881B1B]/10 text-[#881B1B] flex items-center justify-center rounded-full mx-auto">
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-8 h-8 animate-bounce">
							<path stroke-linecap="round" stroke-linejoin="round" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
						</svg>
					</div>
					<div class="space-y-2">
						<h2 class="text-xl font-bold font-serif text-[#0b1535]">{currentTab} Under Construction</h2>
						<p class="text-xs text-slate-500 max-w-md mx-auto leading-relaxed">
							The features for the <strong>{currentTab}</strong> module are scheduled for development by other contributors.
						</p>
					</div>
					<button
						onclick={() => (currentTab = 'Dashboard')}
						class="px-5 py-2.5 bg-[#881B1B] hover:bg-[#881B1B]/95 text-white font-bold text-xs uppercase tracking-wider rounded-lg transition-colors"
					>
						Back to Dashboard
					</button>
				</div>
			{/if}
		</main>
	</div>
</div>
