<script lang="ts">
	// Timeline steps
	const steps = [
		{
			phase: '01',
			title: 'Awareness',
			desc: 'Department briefings and orientation sessions regarding the upcoming cell cycle.'
		},
		{
			phase: '02',
			title: 'Enrollment',
			desc: 'Registration of interest and selection of primary development tracks.'
		},
		{
			phase: '03',
			title: 'Workshops',
			desc: 'Direct instruction sessions in skill tracks and soft-skills seminars.'
		},
		{
			phase: '04',
			title: 'Project Application',
			desc: 'Implementation of learnings in verified, peer-evaluated projects.'
		},
		{
			phase: '05',
			title: 'Assessment',
			desc: 'Committee review of submitted logs, credentials, and achievements.'
		},
		{
			phase: '06',
			title: 'Certification',
			desc: 'Awarding of extra-curricular grades and university-backed certificates.'
		}
	];
</script>

<section class="py-16 sm:py-24 bg-white border-t border-slate-200">
	<div class="max-w-7xl mx-auto px-6 lg:px-8">
		
		<!-- Section Header -->
		<div class="border-b border-slate-100 pb-8 mb-16">
			<span class="text-xs font-bold tracking-widest text-accent-gold uppercase font-sans">
				Activity Process
			</span>
			<h2 class="text-3xl sm:text-4xl font-bold font-serif text-primary-dark tracking-tight mt-2">
				The Student Event Journey
			</h2>
			<p class="text-sm text-slate-500 font-sans max-w-3xl mt-3 leading-relaxed">
				We enforce a structured advancement pipeline to ensure quality and accountability. Below is the official progression timeline for cell members.
			</p>
		</div>

		<!-- Timeline Container -->
		<div class="relative">
			<!-- Timeline path line (Desktop) -->
			<div class="hidden lg:block absolute top-[28px] left-0 right-0 h-0.5 bg-slate-100 -z-10"></div>
			
			<!-- Timeline Grid -->
			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-6 gap-10 lg:gap-6">
				{#each steps as step, i}
					<div class="relative flex flex-col items-start space-y-4">
						<!-- Timeline Connector Node -->
						<div class="flex items-center justify-between w-full lg:w-auto">
							<div class="w-14 h-14 bg-bg-warm border border-slate-200 text-slate-800 font-serif font-bold text-lg flex items-center justify-center rounded-sm shrink-0 shadow-2xs">
								{step.phase}
							</div>
							
							<!-- Arrow indicators for progression (Desktop only, except last) -->
							{#if i < steps.length - 1}
								<div class="hidden lg:block ml-auto text-slate-350">
									<svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
										<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
									</svg>
								</div>
							{/if}
						</div>

						<!-- Content -->
						<div class="space-y-1">
							<h3 class="text-base font-bold text-primary-dark font-sans tracking-tight">
								{step.title}
							</h3>
							<p class="text-xs text-slate-500 leading-relaxed font-sans font-normal">
								{step.desc}
							</p>
						</div>
					</div>
				{/each}
			</div>
		</div>

	</div>
</section>
