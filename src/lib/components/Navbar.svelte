<script lang="ts">
	import { page } from '$app/state';
	import { slide } from 'svelte/transition';

	// Interface for navigation links
	interface NavItem {
		label: string;
		href: string;
	}

	// Props configuration using Svelte 5 runes
	let {
		navItems = [
			{ label: 'Home', href: '/' },
			{ label: 'Tracks', href: '/tracks' },
			{ label: 'Activities', href: '/activities' },
			{ label: 'Leaderboard', href: '/leaderboard' },
			{ label: 'Reports', href: '/reports' },
			{ label: 'About', href: '/about' },
			{ label: 'Contact', href: '/contact' }
		] as NavItem[],
		loginHref = '/login',
		ctaHref = '/get-started',
		ctaText = 'Get Started'
	} = $props();

	// Reactive state for the mobile menu drawer
	let isOpen = $state(false);

	// Toggle state handlers
	function toggleMenu() {
		isOpen = !isOpen;
	}

	function closeMenu() {
		isOpen = false;
	}

	// Helper to check if a navigation item is active
	function isActive(href: string): boolean {
		const path = page.url.pathname;
		if (href === '/') {
			return path === '/';
		}
		return path === href || path.startsWith(href + '/');
	}
</script>

<!-- Sticky wrapper -->
<header class="sticky top-0 z-50 w-full bg-white border-b border-border-gray shadow-sm backdrop-blur-md/95">
	<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div class="flex items-center justify-between h-[72px]">
			
			<!-- Left Section: Logo & Branding -->
			<a href="/" onclick={closeMenu} class="flex items-center gap-3 group focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue focus-visible:ring-offset-2 rounded-xl transition-all duration-200">
				<!-- Custom SVG Logo matching the Figma design -->
				<div class="relative w-10 h-10 flex items-center justify-center rounded-xl shadow-md shadow-primary-blue/10 group-hover:shadow-lg group-hover:shadow-primary-blue/20 transition-all duration-300">
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 40 40" class="w-10 h-10" fill="none">
						<defs>
							<linearGradient id="logo-bg-grad" x1="0%" y1="0%" x2="100%" y2="100%">
								<stop offset="0%" stop-color="#123A84" />
								<stop offset="100%" stop-color="#0A1931" />
							</linearGradient>
							<linearGradient id="spark-grad" x1="0%" y1="0%" x2="100%" y2="100%">
								<stop offset="0%" stop-color="#FFD700" />
								<stop offset="100%" stop-color="#F4B400" />
							</linearGradient>
						</defs>
						<!-- Squircle container -->
						<rect width="40" height="40" rx="12" fill="url(#logo-bg-grad)" />
						<!-- Main 4-pointed Spark Star -->
						<path d="M20 9C20 15 21.5 19 27 20C21.5 21 20 25 20 31C20 25 18.5 21 13 20C18.5 19 20 15 20 9Z" fill="url(#spark-grad)" />
						<!-- Accent Small Star -->
						<path d="M29 11C29 12.5 29.5 13.5 31 14C29.5 14.5 29 15.5 29 17C29 15.5 28.5 14.5 27 14C28.5 13.5 29 12.5 29 11Z" fill="url(#spark-grad)" opacity="0.8" />
						<!-- Spark/Orbit Ring -->
						<path d="M12 28C13 30 17 31 20 30C25 28.5 27.5 22 25 18" stroke="#F4B400" stroke-width="1.2" stroke-linecap="round" opacity="0.5" />
					</svg>
				</div>
				<!-- Brand Text -->
				<div class="flex flex-col justify-center">
					<span class="text-xl font-bold tracking-tight text-primary-blue leading-none transition-colors group-hover:text-primary-blue/90 font-sans">iSPARC</span>
					<span class="text-[9px] font-bold text-gray-400 tracking-wider mt-0.5 uppercase font-sans">IIPS DAVV</span>
				</div>
			</a>

			<!-- Center Section: Desktop Navigation Links -->
			<nav class="hidden lg:flex items-center gap-7" aria-label="Desktop Navigation">
				<ul class="flex items-center gap-7">
					{#each navItems as item}
						<li>
							<a
								href={item.href}
								class="text-[15px] transition-all duration-200 hover:text-primary-blue focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue focus-visible:ring-offset-2 rounded-md px-2 py-1 font-sans {isActive(item.href) ? 'text-primary-blue font-semibold' : 'text-slate-500 font-medium'}"
								aria-current={isActive(item.href) ? 'page' : undefined}
							>
								{item.label}
							</a>
						</li>
					{/each}
				</ul>
			</nav>

			<!-- Right Section: Desktop CTAs -->
			<div class="hidden lg:flex items-center gap-6">
				<a
					href={loginHref}
					class="text-[15px] font-semibold text-primary-blue hover:text-primary-blue/80 transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue focus-visible:ring-offset-2 rounded-md px-2 py-1 font-sans"
				>
					Login
				</a>
				<a
					href={ctaHref}
					class="inline-flex items-center justify-center bg-accent-yellow text-primary-blue font-bold text-[15px] px-6 py-2.5 rounded-lg hover:bg-accent-yellow/95 hover:shadow-md hover:shadow-accent-yellow/20 active:scale-[0.98] transition-all duration-200 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue focus-visible:ring-offset-2 font-sans"
				>
					{ctaText}
				</a>
			</div>

			<!-- Mobile Menu Button (Hamburger) -->
			<div class="flex lg:hidden items-center">
				<button
					onclick={toggleMenu}
					type="button"
					class="relative flex flex-col justify-center items-center w-10 h-10 rounded-lg text-primary-blue hover:bg-slate-50 transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue"
					aria-label={isOpen ? 'Close menu' : 'Open menu'}
					aria-expanded={isOpen}
					aria-controls="mobile-menu"
				>
					<span class="sr-only">Toggle Menu</span>
					<!-- Animated Hamburger Bars -->
					<span class="w-5.5 h-0.5 bg-current rounded-full transition-all duration-300 {isOpen ? 'rotate-45 translate-y-[6px]' : ''}"></span>
					<span class="w-5.5 h-0.5 bg-current rounded-full my-[5px] transition-all duration-300 {isOpen ? 'opacity-0' : ''}"></span>
					<span class="w-5.5 h-0.5 bg-current rounded-full transition-all duration-300 {isOpen ? '-rotate-45 -translate-y-[6px]' : ''}"></span>
				</button>
			</div>

		</div>
	</div>

	<!-- Mobile Dropdown Menu -->
	{#if isOpen}
		<div
			transition:slide={{ duration: 250 }}
			class="lg:hidden border-t border-border-gray bg-white shadow-lg overflow-hidden"
			id="mobile-menu"
		>
			<nav class="px-4 py-5 space-y-4" aria-label="Mobile Navigation">
				<ul class="space-y-1">
					{#each navItems as item}
						<li>
							<a
								href={item.href}
								onclick={closeMenu}
								class="block w-full px-3 py-2.5 rounded-lg text-base font-semibold transition-all duration-200 font-sans {isActive(item.href) ? 'bg-primary-blue/5 text-primary-blue' : 'text-slate-600 hover:bg-slate-50 hover:text-primary-blue'}"
								aria-current={isActive(item.href) ? 'page' : undefined}
							>
								{item.label}
							</a>
						</li>
					{/each}
				</ul>
				
				<!-- Separator -->
				<div class="h-px bg-slate-100 my-4"></div>

				<!-- Mobile CTAs -->
				<div class="flex flex-col gap-3 px-3 pb-2">
					<a
						href={loginHref}
						onclick={closeMenu}
						class="flex items-center justify-center w-full py-2.5 rounded-lg text-base font-bold text-primary-blue hover:bg-slate-50 transition-colors border border-transparent focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue font-sans"
					>
						Login
					</a>
					<a
						href={ctaHref}
						onclick={closeMenu}
						class="flex items-center justify-center w-full py-2.5 rounded-lg text-base font-bold bg-accent-yellow text-primary-blue shadow-sm hover:bg-accent-yellow/95 transition-all duration-200 active:scale-[0.99] focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-blue font-sans"
					>
						{ctaText}
					</a>
				</div>
			</nav>
		</div>
	{/if}
</header>
