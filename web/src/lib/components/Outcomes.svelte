<script lang="ts">
	// Point Allocation Rules list
	const rules = [
		{
			title: 'National / State Representation',
			desc: 'Up to 20 points per event for representing DAVV in national debate matches or athletic tournaments.'
		},
		{
			title: 'College Hackathons & Project Meets',
			desc: '5 to 15 points depending on rank and tier of involvement.'
		},
		{
			title: 'Social Service & Volunteering (NSS/NGO)',
			desc: '1 point per 10 hours of verified field work, up to a maximum of 25 points per semester.'
		}
	];
</script>

<section id="grading-scheme" class="py-16 sm:py-24 bg-bg-warm border-t border-slate-200">
	<div class="max-w-7xl mx-auto px-6 lg:px-8 space-y-16">
		
		<!-- Accreditation and Point Rules grid -->
		<div class="grid grid-cols-1 lg:grid-cols-2 gap-12 lg:gap-16 items-start">
			
			<!-- Left Column: Accreditation & Quote -->
			<div class="space-y-8">
				<h2 class="text-3xl font-bold font-serif text-primary-dark tracking-tight">
					Accreditation & Grading Scheme
				</h2>
				<p class="text-sm text-slate-650 leading-relaxed font-sans font-normal">
					All student certificates and service logs are audited by the departmental committee before final grades are submitted to the university.
				</p>
				
				<!-- Academic Blockquote (Figma spec) -->
				<blockquote class="border-l-2 border-accent-gold pl-6 py-1 bg-white/40 rounded-sm">
					<p class="text-sm italic text-slate-700 font-serif leading-relaxed">
						"Personality development is not merely an addition; it is an essential component of professional readiness."
					</p>
					<cite class="text-[10px] font-bold tracking-wider text-slate-400 uppercase font-sans mt-3 block">
						— DAVV BOARD OF STUDIES GUIDELINES
					</cite>
				</blockquote>
			</div>

			<!-- Right Column: Point Allocation Rules -->
			<div class="bg-white border border-slate-200 p-8 shadow-xs rounded-sm space-y-6">
				<h3 class="text-lg font-bold font-serif text-primary-dark">
					Point Allocation Rules
				</h3>
				
				<ul class="space-y-6">
					{#each rules as rule}
						<li class="flex items-start gap-3">
							<!-- Red Square Bullet (Figma spec) -->
							<span class="w-2.5 h-2.5 bg-[#881B1B] mt-1 shrink-0"></span>
							<div>
								<h4 class="text-sm font-bold text-slate-800 font-sans">{rule.title}:</h4>
								<p class="text-xs text-slate-600 leading-relaxed font-sans font-normal mt-1">{rule.desc}</p>
							</div>
						</li>
					{/each}
				</ul>
			</div>

		</div>

		<!-- Bottom Banner: Student Registry Portal (Dark elegant theme) -->
		<div class="bg-primary-dark text-white p-10 sm:p-14 text-center rounded-sm relative overflow-hidden shadow-sm">
			<!-- Subtle gold accent block line -->
			<div class="absolute top-0 left-0 right-0 h-1.5 bg-accent-gold"></div>
			
			<div class="max-w-2xl mx-auto space-y-6">
				<h2 class="text-2xl sm:text-3xl font-bold font-serif tracking-tight">
					Student Registry Portal
				</h2>
				<p class="text-sm text-slate-350 leading-relaxed font-sans font-normal">
					Access your student profile, upload credentials, view points status, and verify your semester extra-curricular grades.
				</p>
				
				<div class="pt-2">
					<a
						href="/portal"
						class="inline-flex items-center justify-center bg-[#881B1B] hover:bg-[#731717] text-white font-bold text-xs tracking-wider uppercase px-8 py-3.5 border border-[#881B1B] transition-colors duration-250 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent-gold focus-visible:ring-offset-2"
					>
						Enter Portal
					</a>
				</div>
			</div>
		</div>

	</div>
</section>
