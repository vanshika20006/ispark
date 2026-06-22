<script lang="ts">
	// Track structures
	const personalityTrack = {
		title: 'Personality Development Track',
		desc: 'Designed to elevate students\' interpersonal qualities, leadership traits, and professional presentation capabilities to align with modern corporate and research expectations.',
		modules: [
			{ name: 'Grooming & Presentation', focus: 'Professional attire, body language, and corporate etiquette.' },
			{ name: 'Effective Communication', focus: 'Public speaking, active listening, and business writing.' },
			{ name: 'Teamwork & Collaboration', focus: 'Group dynamics, conflict resolution, and peer leadership.' },
			{ name: 'GD & Interview Readiness', focus: 'Mock group discussions, resume crafting, and technical interview simulations.' }
		]
	};

	const skillTrack = {
		title: 'Skill Building Track',
		desc: 'Hands-on technical advancement pathways structured around high-demand industry skills, software systems, and engineering frameworks.',
		modules: [
			{ name: 'Artificial Intelligence', focus: 'Fundamentals of machine learning, neural networks, and prompt engineering.' },
			{ name: 'Data Analytics & Stats', focus: 'Data visualization, SQL, python analytics, and business intelligence.' },
			{ name: 'Cyber Security Operations', focus: 'Network security, secure coding, threat intelligence, and ethical hacking.' },
			{ name: 'Project Management & Digital', focus: 'Agile methodologies, digital marketing campaigns, and systems administration.' }
		]
	};
</script>

<section id="tracks" class="py-16 sm:py-24 bg-bg-warm border-t border-slate-200">
	<div class="max-w-7xl mx-auto px-6 lg:px-8">
		
		<!-- Section Header -->
		<div class="text-center max-w-3xl mx-auto mb-16 space-y-4">
			<span class="text-xs font-bold tracking-widest text-accent-gold uppercase font-sans">
				Academic Curriculum
			</span>
			<h2 class="text-3xl sm:text-4xl font-bold font-serif text-primary-dark tracking-tight">
				Student Advancement Tracks
			</h2>
			<div class="h-0.5 w-12 bg-accent-gold mx-auto mt-4"></div>
			<p class="text-sm sm:text-base text-slate-500 font-sans leading-relaxed">
				ISPARC facilitates two core learning pipelines tailored to cultivate well-rounded professionals. These paths combine technical mastery with essential leadership and behavioral grooming.
			</p>
		</div>

		<!-- Tracks Grid (2 Columns) -->
		<div class="grid grid-cols-1 lg:grid-cols-2 gap-10 lg:gap-12">
			
			<!-- Track 1: Personality Development -->
			<div class="bg-white border border-slate-200 p-8 sm:p-10 rounded-sm flex flex-col justify-between shadow-xs">
				<div class="space-y-6">
					<div class="inline-flex items-center px-2.5 py-1 bg-amber-50 text-[10px] font-bold text-accent-gold border border-amber-200 uppercase tracking-wider">
						Track A
					</div>
					<h3 class="text-2xl font-bold font-serif text-primary-dark tracking-tight">
						{personalityTrack.title}
					</h3>
					<p class="text-sm text-slate-650 leading-relaxed font-sans font-normal">
						{personalityTrack.desc}
					</p>
					
					<!-- Syllabus List -->
					<div class="space-y-4 border-t border-slate-100 pt-6">
						<h4 class="text-xs font-bold text-slate-400 uppercase tracking-widest">Syllabus Highlights</h4>
						<ul class="space-y-4">
							{#each personalityTrack.modules as mod}
								<li class="flex items-start gap-3">
									<span class="w-1.5 h-1.5 bg-accent-gold mt-1.5 shrink-0"></span>
									<div>
										<span class="text-sm font-semibold text-slate-800 block">{mod.name}</span>
										<span class="text-xs text-slate-550 leading-relaxed block mt-0.5">{mod.focus}</span>
									</div>
								</li>
							{/each}
						</ul>
					</div>
				</div>
			</div>

			<!-- Track 2: Skill Building -->
			<div class="bg-white border border-slate-200 p-8 sm:p-10 rounded-sm flex flex-col justify-between shadow-xs">
				<div class="space-y-6">
					<div class="inline-flex items-center px-2.5 py-1 bg-amber-50 text-[10px] font-bold text-accent-gold border border-amber-200 uppercase tracking-wider">
						Track B
					</div>
					<h3 class="text-2xl font-bold font-serif text-primary-dark tracking-tight">
						{skillTrack.title}
					</h3>
					<p class="text-sm text-slate-650 leading-relaxed font-sans font-normal">
						{skillTrack.desc}
					</p>
					
					<!-- Syllabus List -->
					<div class="space-y-4 border-t border-slate-100 pt-6">
						<h4 class="text-xs font-bold text-slate-400 uppercase tracking-widest">Syllabus Highlights</h4>
						<ul class="space-y-4">
							{#each skillTrack.modules as mod}
								<li class="flex items-start gap-3">
									<span class="w-1.5 h-1.5 bg-accent-gold mt-1.5 shrink-0"></span>
									<div>
										<span class="text-sm font-semibold text-slate-800 block">{mod.name}</span>
										<span class="text-xs text-slate-550 leading-relaxed block mt-0.5">{mod.focus}</span>
									</div>
								</li>
							{/each}
						</ul>
					</div>
				</div>
			</div>

		</div>

	</div>
</section>
