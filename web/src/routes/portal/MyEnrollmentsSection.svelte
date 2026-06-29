<script lang="ts">
	

	// Props using Svelte 5 runes
	let { onUploadCertificateClick = () => {} } = $props();

	// KPI Cards data
	const kpis = [
		{
			title: 'Total Enrollments',
			value: 12,
			icon: `M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01`,
			bgColor: 'bg-indigo-50/50',
			iconColor: 'text-indigo-650'
		},
		{
			title: 'Upcoming Activities',
			value: 4,
			icon: `M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z`,
			bgColor: 'bg-blue-50/50',
			iconColor: 'text-blue-650'
		},
		{
			title: 'Completed Activities',
			value: 6,
			icon: `M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z`,
			bgColor: 'bg-emerald-50/50',
			iconColor: 'text-emerald-600'
		},
		{
			title: 'Pending Verification',
			value: 2,
			icon: `M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z`,
			bgColor: 'bg-amber-50/50',
			iconColor: 'text-amber-600'
		}
	];

	// Registered Activities list
	const registeredActivities = [
		{
			name: 'National Hackathon 2026',
			category: 'Technical',
			enrollDate: '10 Jun 2026',
			activityDate: '05 Jul 2026',
			credits: 15,
			status: 'Registered',
			action: 'details'
		},
		{
			name: 'Inter College Debate',
			category: 'Public Speaking',
			enrollDate: '05 Jun 2026',
			activityDate: '20 Jun 2026',
			credits: 10,
			status: 'Completed',
			action: 'certificate'
		},
		{
			name: 'Blood Donation Drive',
			category: 'Social Service',
			enrollDate: '15 May 2026',
			activityDate: '25 May 2026',
			credits: 5,
			status: 'Pending Verification',
			action: 'upload'
		},
		{
			name: 'Robotics Workshop',
			category: 'Technical',
			enrollDate: '02 Apr 2026',
			activityDate: '18 Apr 2026',
			credits: 12,
			status: 'Verified',
			action: 'certificate'
		},
		{
			name: 'Annual Cultural Fest',
			category: 'Cultural',
			enrollDate: '01 Mar 2026',
			activityDate: '20 Mar 2026',
			credits: 8,
			status: 'Completed',
			action: 'certificate'
		},
		{
			name: 'Research Presentation Contest',
			category: 'Academic',
			enrollDate: '15 Feb 2026',
			activityDate: '10 Mar 2026',
			credits: 12,
			status: 'Verified',
			action: 'certificate'
		},
		{
			name: 'National Science Olympiad',
			category: 'Technical',
			enrollDate: '20 Jan 2026',
			activityDate: '05 Feb 2026',
			credits: 10,
			status: 'Rejected',
			action: 'details'
		}
	];

	// Selected activity for participation journey tracking
	let selectedJourneyActivity = $state('Inter College Debate');

	// Map activity name to journey steps
	const journeySteps = [
		{ title: 'Enrolled', desc: 'Registered for the activity', state: 'completed' },
		{ title: 'Activity Started', desc: 'Event began on scheduled date', state: 'completed' },
		{ title: 'Completed', desc: 'Participated in all sessions', state: 'completed' },
		{ title: 'Certificate Submitted', desc: 'Uploaded proof of participation', state: 'current' },
		{ title: 'Verified', desc: 'Awaiting mentor verification', state: 'pending' },
		{ title: 'Credits Awarded', desc: 'Credits will reflect in profile', state: 'pending' }
	];

	// Upcoming Activities
	const upcomingActivities = [
		{
			name: 'National Hackathon 2026',
			date: '5 July 2026',
			time: '10:00 AM',
			location: 'IIPS Auditorium',
			credits: 15,
			timeLeft: '12d left'
		},
		{
			name: 'Innovation Challenge',
			date: '12 July 2026',
			time: '09:00 AM',
			location: 'Seminar Hall B',
			credits: 10,
			timeLeft: '19d left'
		},
		{
			name: 'Technical Workshop Series',
			date: '20 July 2026',
			time: '11:00 AM',
			location: 'Computer Lab 3',
			credits: 8,
			timeLeft: '27d left'
		},
		{
			name: 'Research Symposium',
			date: '2 Aug 2026',
			time: '10:30 AM',
			location: 'Conference Room',
			credits: 12,
			timeLeft: '40d left'
		}
	];

	// Enrollment Credit Summary progress categories
	const creditCategories = [
		{ name: 'Technical Activities', credits: 48, max: 60, color: 'bg-[#881B1B]' },
		{ name: 'Public Speaking', credits: 24, max: 40, color: 'bg-[#0B1535]' },
		{ name: 'Social Service', credits: 18, max: 40, color: 'bg-emerald-600' },
		{ name: 'Sports', credits: 12, max: 40, color: 'bg-blue-600' }
	];

	// Recent updates
	const recentUpdates = [
		{
			text: 'National Hackathon registration confirmed',
			time: '2 days ago',
			icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z',
			color: 'text-emerald-600 bg-emerald-50'
		},
		{
			text: '5 credits added for Robotics Workshop',
			time: '5 days ago',
			icon: 'M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.907c.969 0 1.371 1.24.588 1.81l-3.97 2.883a1 1 0 00-.364 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.971-2.883a1 1 0 00-1.18 0l-3.97 2.883c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.364-1.118l-3.97-2.883c-.783-.57-.38-1.81.588-1.81h4.906a1 1 0 00.95-.69l1.519-4.674z',
			color: 'text-indigo-600 bg-indigo-50'
		},
		{
			text: 'Certificate verified for Annual Fest',
			time: '1 week ago',
			icon: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z',
			color: 'text-teal-600 bg-teal-50'
		},
		{
			text: 'Enrolled in Innovation Challenge',
			time: '2 weeks ago',
			icon: 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253',
			color: 'text-blue-600 bg-blue-50'
		}
	];

	function getStatusClass(status: string) {
		switch (status) {
			case 'Registered':
				return 'bg-blue-50 text-blue-700 border-blue-150';
			case 'Completed':
				return 'bg-emerald-50 text-emerald-700 border-emerald-150';
			case 'Verified':
				return 'bg-emerald-50 text-emerald-800 border-emerald-250 font-bold';
			case 'Pending Verification':
				return 'bg-amber-50 text-amber-850 border-amber-250';
			case 'Rejected':
				return 'bg-rose-50 text-rose-700 border-rose-150';
			default:
				return 'bg-slate-50 text-slate-700 border-slate-150';
		}
	}
</script>

<div class="space-y-6 w-full animate-fade-in font-sans">
	<!-- KPI Grid -->
	<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
		{#each kpis as kpi}
			<div
				class="bg-white border border-slate-200 p-5 rounded-xl flex items-center gap-4 hover:shadow-xs transition duration-200"
			>
				<div class="w-12 h-12 rounded-lg {kpi.bgColor} flex items-center justify-center shrink-0">
					<svg
						class="w-6 h-6 {kpi.iconColor}"
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
						stroke-width="2"
					>
						<path stroke-linecap="round" stroke-linejoin="round" d={kpi.icon} />
					</svg>
				</div>
				<div class="flex flex-col">
					<span class="text-2xl font-extrabold text-[#0B1535] leading-none">{kpi.value}</span>
					<span class="text-xs font-semibold text-[#6B7280] mt-1.5">{kpi.title}</span>
				</div>
			</div>
		{/each}
	</div>

	<!-- Table Card -->
	<div class="bg-white border border-slate-200 rounded-xl overflow-hidden shadow-xs">
		<div class="p-5 border-b border-slate-100 flex items-center justify-between">
			<div>
				<h2 class="text-base font-bold font-serif text-[#0B1535]">My Registered Activities</h2>
				<p class="text-[10px] font-semibold text-[#6B7280] uppercase tracking-widest mt-0.5">
					Academic Year 2025 - 2026
				</p>
			</div>
			<span
				class="text-xs font-semibold text-slate-500 bg-slate-50 border border-slate-150 px-2.5 py-1 rounded-full"
			>
				7 records
			</span>
		</div>

		<div class="overflow-x-auto">
			<table class="w-full text-left border-collapse">
				<thead>
					<tr
						class="bg-slate-50 border-b border-slate-150 text-[10px] font-bold text-slate-400 uppercase tracking-wider"
					>
						<th class="px-6 py-4">Activity Name</th>
						<th class="px-6 py-4">Category</th>
						<th class="px-6 py-4">Enrollment Date</th>
						<th class="px-6 py-4">Activity Date</th>
						<th class="px-6 py-4 text-center">Credits</th>
						<th class="px-6 py-4">Status</th>
						<th class="px-6 py-4 text-right">Actions</th>
					</tr>
				</thead>
				<tbody class="divide-y divide-slate-100 text-[13px]">
					{#each registeredActivities as act}
						<tr class="hover:bg-slate-50/50 transition-colors duration-150">
							<td class="px-6 py-4 font-bold text-[#0B1535]">{act.name}</td>
							<td class="px-6 py-4 text-slate-500 font-medium">{act.category}</td>
							<td class="px-6 py-4 text-slate-550">{act.enrollDate}</td>
							<td class="px-6 py-4 text-slate-550">{act.activityDate}</td>
							<td class="px-6 py-4 text-center font-extrabold text-[#881B1B]">{act.credits}</td>
							<td class="px-6 py-4">
								<span
									class="inline-flex items-center px-2.5 py-0.5 rounded-full text-[10px] font-bold tracking-wide border {getStatusClass(
										act.status
									)}"
								>
									{act.status}
								</span>
							</td>
							<td class="px-6 py-4 text-right">
								{#if act.action === 'details'}
									<button
										class="inline-flex items-center justify-center border border-slate-250 hover:border-slate-800 text-slate-705 hover:bg-slate-900 hover:text-white px-3.5 py-2 rounded-lg text-xs font-bold transition duration-150 cursor-pointer"
									>
										<svg
											class="w-3.5 h-3.5 mr-1.5"
											fill="none"
											stroke="currentColor"
											stroke-width="2"
											viewBox="0 0 24 24"
										>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
											/>
											<path
												stroke-linecap="round"
												stroke-linejoin="round"
												d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
											/>
										</svg>
										View Details
									</button>
								{:else}
									<button
										onclick={() => {
											if (act.action === 'upload') {
												onUploadCertificateClick();
											} else {
												alert(`Download/View Certificate for ${act.name}`);
											}
										}}
										class="inline-flex items-center justify-center px-3.5 py-2 rounded-lg text-xs font-bold transition duration-150 cursor-pointer {act.action ===
										'upload'
											? 'bg-[#881B1B] hover:bg-[#731717] text-white shadow-xs'
											: 'border border-slate-250 hover:border-slate-800 text-slate-705 hover:bg-slate-900 hover:text-white'}"
									>
										{#if act.action === 'upload'}
											<svg
												class="w-3.5 h-3.5 mr-1.5"
												fill="none"
												stroke="currentColor"
												stroke-width="2.5"
												viewBox="0 0 24 24"
											>
												<path
													stroke-linecap="round"
													stroke-linejoin="round"
													d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
												/>
											</svg>
											Upload Certificate
										{:else}
											<svg
												class="w-3.5 h-3.5 mr-1.5"
												fill="none"
												stroke="currentColor"
												stroke-width="2"
												viewBox="0 0 24 24"
											>
												<path
													stroke-linecap="round"
													stroke-linejoin="round"
													d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"
												/>
											</svg>
											View Certificate
										{/if}
									</button>
								{/if}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>

	<!-- Two-column workspace -->
	<div class="grid grid-cols-1 lg:grid-cols-12 gap-6">
		<!-- Left Column -->
		<div class="lg:col-span-7 space-y-6">
			<!-- Participation Journey -->
			<div class="bg-white border border-slate-200 p-6 rounded-xl shadow-xs">
				<div class="border-b border-slate-100 pb-4 mb-5">
					<h2 class="text-base font-bold font-serif text-[#0B1535]">Participation Journey</h2>
					<p class="text-[10px] font-bold text-slate-400 uppercase tracking-widest mt-0.5">
						{selectedJourneyActivity} • Tracking Status
					</p>
				</div>

				<div class="relative pl-6 border-l-2 border-slate-150 ml-3 space-y-6 py-2">
					{#each journeySteps as step}
						<div class="relative">
							<!-- Timeline Indicator Dot -->
							<div
								class="absolute -left-[31px] top-1 w-4 h-4 rounded-full border-2 bg-white flex items-center justify-center transition duration-200
								{step.state === 'completed' ? 'border-[#0B1535] bg-[#0B1535]' : ''}
								{step.state === 'current' ? 'border-[#881B1B] bg-white ring-4 ring-[#881B1B]/10' : ''}
								{step.state === 'pending' ? 'border-slate-200' : ''}
							"
							>
								{#if step.state === 'completed'}
									<svg
										class="w-2.5 h-2.5 text-white"
										fill="none"
										stroke="currentColor"
										stroke-width="3"
										viewBox="0 0 24 24"
									>
										<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
									</svg>
								{:else if step.state === 'current'}
									<div class="w-1.5 h-1.5 rounded-full bg-[#881B1B]"></div>
								{/if}
							</div>

							<!-- Timeline Text -->
							<div class="flex items-center justify-between gap-4">
								<div>
									<h4
										class="text-xs font-bold {step.state === 'pending'
											? 'text-slate-400'
											: 'text-[#0B1535]'}"
									>
										{step.title}
									</h4>
									<p class="text-[11px] text-[#6B7280] mt-0.5 leading-relaxed">{step.desc}</p>
								</div>

								{#if step.state === 'current'}
									<span
										class="bg-amber-50 text-amber-800 border border-amber-250 px-2 py-0.5 rounded-md text-[9px] font-bold uppercase tracking-wider"
									>
										Current
									</span>
								{/if}
							</div>
						</div>
					{/each}
				</div>
			</div>

			<!-- Upcoming Activities list -->
			<div class="bg-white border border-slate-200 p-6 rounded-xl shadow-xs">
				<div class="border-b border-slate-100 pb-4 mb-5">
					<h2 class="text-base font-bold font-serif text-[#0B1535]">Upcoming Activities</h2>
					<p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest mt-0.5">
						Enrolled and verified schedule
					</p>
				</div>

				<div class="divide-y divide-slate-100">
					{#each upcomingActivities as act}
						<div class="flex items-start justify-between py-4 first:pt-0 last:pb-0 gap-4">
							<div class="flex items-start gap-4">
								<div
									class="w-10 h-10 rounded-lg bg-slate-50 border border-slate-150 flex items-center justify-center shrink-0"
								>
									<svg
										class="w-5 h-5 text-slate-450"
										xmlns="http://www.w3.org/2000/svg"
										fill="none"
										viewBox="0 0 24 24"
										stroke="currentColor"
										stroke-width="2"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
										/>
									</svg>
								</div>
								<div class="space-y-1">
									<h4 class="text-xs font-bold text-[#0B1535]">{act.name}</h4>
									<div
										class="flex flex-wrap items-center gap-x-2.5 gap-y-1 text-[10px] text-[#6B7280]"
									>
										<span class="flex items-center">
											<svg
												class="w-3.5 h-3.5 mr-1"
												fill="none"
												stroke="currentColor"
												viewBox="0 0 24 24"
												><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
												/></svg
											>
											{act.date}
										</span>
										<span class="flex items-center">
											<svg
												class="w-3.5 h-3.5 mr-1"
												fill="none"
												stroke="currentColor"
												viewBox="0 0 24 24"
												><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
												/></svg
											>
											{act.time}
										</span>
										<span class="flex items-center">
											<svg
												class="w-3.5 h-3.5 mr-1"
												fill="none"
												stroke="currentColor"
												viewBox="0 0 24 24"
												><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
												/><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
												/></svg
											>
											{act.location}
										</span>
									</div>
								</div>
							</div>

							<div class="text-right flex flex-col items-end shrink-0">
								<span class="text-xs font-extrabold text-[#881B1B]">{act.credits} Credits</span>
								<span
									class="text-[9px] font-bold text-amber-700 bg-amber-50 border border-amber-150 px-1.5 py-0.5 rounded-md mt-1.5 uppercase tracking-wide"
								>
									{act.timeLeft}
								</span>
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>

		<!-- Right Column -->
		<div class="lg:col-span-5 space-y-6">
			<!-- Enrollment Credit Summary -->
			<div class="bg-white border border-slate-200 p-6 rounded-xl shadow-xs">
				<div class="border-b border-slate-100 pb-4 mb-5">
					<h2 class="text-base font-bold font-serif text-[#0B1535]">Enrollment Credit Summary</h2>
					<p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest mt-0.5">
						Credit milestones breakdown
					</p>
				</div>

				<div class="space-y-4">
					{#each creditCategories as cat}
						<div class="space-y-1.5">
							<div class="flex items-center justify-between text-xs">
								<span class="font-bold text-[#0B1535]">{cat.name}</span>
								<span class="font-extrabold text-[#0B1535]">{cat.credits} Credits</span>
							</div>
							<div class="w-full bg-slate-100 h-2 rounded-full overflow-hidden">
								<div
									class="{cat.color} h-2 rounded-full transition-all duration-500"
									style="width: {(cat.credits / cat.max) * 100}%"
								></div>
							</div>
						</div>
					{/each}

					<!-- Total credits bar -->
					<div class="pt-5 border-t border-slate-100 space-y-3">
						<div class="flex items-center justify-between">
							<span class="text-xs font-bold text-slate-600">Total Credits vs Required</span>
							<span class="text-base font-extrabold text-[#881B1B]">118 / 200</span>
						</div>
						<div class="w-full bg-slate-100 h-3 rounded-full overflow-hidden">
							<div
								class="bg-[#881B1B] h-3 rounded-full transition-all duration-500"
								style="width: 59%"
							></div>
						</div>
						<div class="text-center">
							<span class="text-[10px] font-bold text-slate-400 uppercase tracking-wider"
								>59% Complete</span
							>
						</div>
					</div>
				</div>
			</div>

			<!-- Recent Enrollment Updates -->
			<div class="bg-white border border-slate-200 p-6 rounded-xl shadow-xs">
				<div class="border-b border-slate-100 pb-4 mb-5">
					<h2 class="text-base font-bold font-serif text-[#0B1535]">Recent Enrollment Updates</h2>
					<p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest mt-0.5">
						Activity timeline logs
					</p>
				</div>

				<div class="space-y-4.5">
					{#each recentUpdates as update}
						<div class="flex items-start gap-3.5">
							<div
								class="w-7.5 h-7.5 rounded-full {update.color} flex items-center justify-center shrink-0"
							>
								<svg
									class="w-4.5 h-4.5"
									fill="none"
									stroke="currentColor"
									stroke-width="2"
									viewBox="0 0 24 24"
								>
									<path stroke-linecap="round" stroke-linejoin="round" d={update.icon} />
								</svg>
							</div>
							<div class="space-y-0.5">
								<p class="text-xs font-semibold text-[#0B1535] leading-relaxed">{update.text}</p>
								<span class="text-[10px] text-slate-450 block">{update.time}</span>
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.animate-fade-in {
		opacity: 0;
		animation: fadeIn 0.4s ease-out forwards;
	}

	@keyframes fadeIn {
		to {
			opacity: 1;
		}
	}
</style>
