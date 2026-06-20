<script lang="ts">
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<div class="min-h-screen bg-slate-50 flex flex-col font-sans">
	<!-- Navbar mount -->
	<Navbar />

	<!-- Hero section -->
	<main class="flex-grow flex items-center justify-center py-20 px-4 sm:px-6 lg:px-8 bg-linear-to-b from-white to-slate-50">
		<div class="max-w-4xl text-center space-y-8">
			
			<!-- Animated Badge -->
			<div class="inline-flex items-center gap-2 px-4 py-1.5 bg-amber-50 border border-amber-200 rounded-full text-amber-800 text-xs font-semibold tracking-wide shadow-xs animate-bounce">
				<svg class="w-3.5 h-3.5 text-amber-500" fill="currentColor" viewBox="0 0 20 20">
					<path fill-rule="evenodd" d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" clip-rule="evenodd" />
				</svg>
				Future Ready Professionals
			</div>

			<!-- Main Heading -->
			<h1 class="text-4xl sm:text-6xl font-extrabold text-primary-blue tracking-tight leading-tight">
				Transforming Skills,<br />
				<span class="bg-linear-to-r from-accent-yellow to-amber-500 bg-clip-text text-transparent">Advancing Personalities</span>
			</h1>

			<!-- Subtitle -->
			<p class="max-w-2xl mx-auto text-lg sm:text-xl text-slate-600 font-medium leading-relaxed">
				Welcome to IIPS Skill, Personality Advancement & Refinement Cell (ISPARC). We empower students at IIPS DAVV to discover their potential and achieve excellence.
			</p>

			<!-- Action buttons -->
			<div class="flex flex-col sm:flex-row items-center justify-center gap-4 pt-4">
				<a
					href="/get-started"
					class="w-full sm:w-auto inline-flex items-center justify-center px-8 py-3.5 rounded-xl bg-accent-yellow text-primary-blue font-extrabold text-base shadow-md shadow-accent-yellow/20 transition-all hover:bg-accent-yellow/95 hover:shadow-lg hover:shadow-accent-yellow/30 active:scale-98 focus:outline-none focus:ring-2 focus:ring-primary-blue focus:ring-offset-2"
				>
					Explore Tracks
				</a>
				<a
					href="/about"
					class="w-full sm:w-auto inline-flex items-center justify-center px-8 py-3.5 rounded-xl bg-white text-primary-blue font-bold text-base border border-slate-200 hover:bg-slate-50 transition-all active:scale-98 focus:outline-none focus:ring-2 focus:ring-primary-blue focus:ring-offset-2"
				>
					Learn More
				</a>
			</div>
		</div>
	</main>

	<!-- Feature Grid to allow scrolling and testing sticky behavior -->
	<section class="py-24 bg-white border-t border-slate-100">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="grid grid-cols-1 md:grid-cols-3 gap-8">
				
				<!-- Card 1 -->
				<div class="p-8 bg-slate-50 rounded-2xl border border-slate-100 hover:border-primary-blue/20 hover:shadow-lg transition-all duration-300 space-y-4 group">
					<div class="w-12 h-12 bg-primary-blue/5 text-primary-blue rounded-xl flex items-center justify-center font-bold text-xl group-hover:bg-primary-blue group-hover:text-white transition-all duration-300">
						01
					</div>
					<h3 class="text-xl font-bold text-primary-blue">Skill Tracks</h3>
					<p class="text-slate-600 leading-relaxed">Structured learning paths covering modern technologies, design systems, and soft skills.</p>
				</div>
				
				<!-- Card 2 -->
				<div class="p-8 bg-slate-50 rounded-2xl border border-slate-100 hover:border-primary-blue/20 hover:shadow-lg transition-all duration-300 space-y-4 group">
					<div class="w-12 h-12 bg-primary-blue/5 text-primary-blue rounded-xl flex items-center justify-center font-bold text-xl group-hover:bg-primary-blue group-hover:text-white transition-all duration-300">
						02
					</div>
					<h3 class="text-xl font-bold text-primary-blue">Cell Activities</h3>
					<p class="text-slate-600 leading-relaxed">Workshops, bootcamps, and coding contests designed to refine students' capabilities.</p>
				</div>
				
				<!-- Card 3 -->
				<div class="p-8 bg-slate-50 rounded-2xl border border-slate-100 hover:border-primary-blue/20 hover:shadow-lg transition-all duration-300 space-y-4 group">
					<div class="w-12 h-12 bg-primary-blue/5 text-primary-blue rounded-xl flex items-center justify-center font-bold text-xl group-hover:bg-primary-blue group-hover:text-white transition-all duration-300">
						03
					</div>
					<h3 class="text-xl font-bold text-primary-blue">Leaderboard</h3>
					<p class="text-slate-600 leading-relaxed">Track your progress, earn badges, and compete with peers on the official ISPARC leaderboard.</p>
				</div>

			</div>
		</div>
	</section>
</div>
