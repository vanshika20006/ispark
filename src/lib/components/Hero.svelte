<script lang="ts">
	// Props configuration
	let {
		badgeText = 'Departmental Affiliated Cell',
		titleText = 'Skill, Personality Advancement & Refinement Cell.',
		descriptionText = 'Welcome to ISPARC, an official cell of the Institute of Infrastructure and Professional Studies (IIPS), DAVV. We facilitate structured extracurricular activities, personality enhancement tracks, and grading systems to foster student growth and leadership.',
		primaryCtaText = 'Grading Matrix',
		primaryCtaHref = '#grading-scheme',
		secondaryCtaText = 'Explore Tracks',
		secondaryCtaHref = '#tracks'
	} = $props();

	// Grading matrix data for the right-hand panel
	const grades = [
		{ range: '60+ Points', grade: 'Grade O (Outstanding)', colorClass: 'bg-emerald-50 text-emerald-800 border-emerald-200' },
		{ range: '45 - 59 Points', grade: 'Grade A (Excellent)', colorClass: 'bg-blue-50 text-blue-800 border-blue-200' },
		{ range: '30 - 44 Points', grade: 'Grade B (Very Good)', colorClass: 'bg-amber-50 text-amber-850 border-amber-200' },
		{ range: '15 - 29 Points', grade: 'Grade C (Good)', colorClass: 'bg-purple-50 text-purple-800 border-purple-200' },
		{ range: 'Under 15 Points', grade: 'Grade D (Satisfactory)', colorClass: 'bg-slate-50 text-slate-800 border-slate-200' }
	];
</script>

<section class="relative py-16 sm:py-24 bg-bg-warm">
	<div class="max-w-7xl mx-auto px-6 lg:px-8">
		<div class="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-16 items-start">
			
			<!-- Left Column: Headline and Context -->
			<div class="lg:col-span-7 space-y-8">
				
				<!-- Academic Affiliation Badge -->
				<div class="inline-flex items-center px-3.5 py-1.5 border border-slate-300 rounded-sm text-[10px] font-bold tracking-widest uppercase text-slate-600 bg-white/50">
					{badgeText}
				</div>

				<!-- Main Institutional Headline -->
				<h1 class="text-4xl sm:text-5xl lg:text-[56px] lg:leading-[1.1] font-bold text-primary-dark font-serif tracking-tight">
					{titleText}
				</h1>

				<!-- Subtext explanations -->
				<p class="text-base sm:text-lg text-slate-650 leading-relaxed font-sans font-normal max-w-2xl">
					{descriptionText}
				</p>

				<!-- Buttons (Sharp, premium editorial design style) -->
				<div class="flex flex-col sm:flex-row gap-4 pt-4">
					<a
						href={primaryCtaHref}
						class="inline-flex items-center justify-center bg-primary-dark text-white font-bold text-xs tracking-wider uppercase px-7 py-4 border border-primary-dark hover:bg-slate-800 transition-colors duration-250 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent-gold"
					>
						{primaryCtaText}
					</a>
					<a
						href={secondaryCtaHref}
						class="inline-flex items-center justify-center border border-slate-300 bg-white/30 text-primary-dark font-bold text-xs tracking-wider uppercase px-7 py-4 hover:border-slate-850 transition-colors duration-250 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent-gold"
					>
						{secondaryCtaText}
					</a>
				</div>

			</div>

			<!-- Right Column: Activity Grading Matrix (Clean Grid Panel) -->
			<div class="lg:col-span-5 bg-white border border-slate-200 p-8 shadow-xs rounded-sm">
				<h2 class="text-xl font-bold font-serif text-primary-dark mb-3">
					Activity Grading Matrix
				</h2>
				<p class="text-xs text-slate-500 font-sans leading-relaxed mb-6 border-b border-slate-100 pb-4">
					Extracurricular points are mapped directly to grades in semester results under the DAVV guidelines.
				</p>
				
				<div class="space-y-3.5">
					{#each grades as item}
						<div class="flex items-center justify-between py-1.5 border-b border-slate-50 last:border-b-0">
							<span class="text-sm font-semibold text-slate-700 font-sans">{item.range}</span>
							<span class="inline-flex items-center px-3 py-1 rounded-sm text-xs font-bold border {item.colorClass} tracking-wide font-sans">
								{item.grade}
							</span>
						</div>
					{/each}
				</div>
			</div>

		</div>
	</div>
</section>
