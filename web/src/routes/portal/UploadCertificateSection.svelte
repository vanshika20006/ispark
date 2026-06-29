<script lang="ts">
	import { fade, slide } from 'svelte/transition';

	// Form field states (Svelte 5 runes)
	let activityName = $state('');
	let activityCategory = $state('');
	let activityDate = $state('');
	let organizerName = $state('');
	let eventLevel = $state('');
	let certificateNumber = $state('');
	let issueDate = $state('');
	let participationType = $state('');
	let achievementDescription = $state('');

	// File upload states
	let selectedFile = $state<File | null>(null);
	let isDragOver = $state(false);
	let submitting = $state(false);
	let submitSuccess = $state(false);
	let errorMsg = $state('');

	// Dropdown options
	const eventLevels = [
		'Department Level',
		'Institute Level',
		'State/University Level',
		'National Level',
		'International Level'
	];

	const participationTypes = [
		'Winner',
		'Runner Up',
		'Active Participant',
		'Organizer / Coordinator',
		'Volunteer Service'
	];

	// Guidelines list
	const guidelines = [
		'Upload clear and readable certificates',
		'Ensure your name is visible',
		'Certificates must be authentic',
		'Supported formats: PDF, JPG, PNG',
		'File size must not exceed 10 MB',
		'One certificate per submission'
	];

	// Automatically map category based on activity keywords (for autofill mock)
	$effect(() => {
		const name = activityName.toLowerCase();
		if (
			name.includes('hackathon') ||
			name.includes('coding') ||
			name.includes('robotics') ||
			name.includes('workshop')
		) {
			activityCategory = 'Technical';
		} else if (
			name.includes('debate') ||
			name.includes('elocution') ||
			name.includes('model united nations') ||
			name.includes('mun')
		) {
			activityCategory = 'Public Speaking';
		} else if (
			name.includes('blood') ||
			name.includes('camp') ||
			name.includes('social') ||
			name.includes('donation') ||
			name.includes('nss')
		) {
			activityCategory = 'Social Service';
		} else if (
			name.includes('sport') ||
			name.includes('cricket') ||
			name.includes('football') ||
			name.includes('tournament')
		) {
			activityCategory = 'Sports';
		} else if (
			name.includes('dance') ||
			name.includes('music') ||
			name.includes('fest') ||
			name.includes('drama')
		) {
			activityCategory = 'Cultural';
		} else if (activityName.trim() !== '') {
			activityCategory = 'General Extracurricular';
		} else {
			activityCategory = '';
		}
	});

	// Handle file selection via input click
	function handleFileChange(event: Event) {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			validateAndSetFile(target.files[0]);
		}
	}

	// Drag and drop event handlers
	function handleDragOver(event: DragEvent) {
		event.preventDefault();
		isDragOver = true;
	}

	function handleDragLeave() {
		isDragOver = false;
	}

	function handleDrop(event: DragEvent) {
		event.preventDefault();
		isDragOver = false;
		if (event.dataTransfer && event.dataTransfer.files && event.dataTransfer.files[0]) {
			validateAndSetFile(event.dataTransfer.files[0]);
		}
	}

	function validateAndSetFile(file: File) {
		errorMsg = '';
		const allowedTypes = ['application/pdf', 'image/jpeg', 'image/png', 'image/jpg'];
		const maxSize = 10 * 1024 * 1024; // 10MB

		if (!allowedTypes.includes(file.type)) {
			errorMsg = 'Invalid file type. Please upload a PDF, PNG, or JPG file.';
			return;
		}

		if (file.size > maxSize) {
			errorMsg = 'File size exceeds the 10 MB limit.';
			return;
		}

		selectedFile = file;
	}

	function removeSelectedFile() {
		selectedFile = null;
	}

	// Submit handler
	function handleSubmit(event: SubmitEvent) {
		event.preventDefault();
		if (!selectedFile) {
			errorMsg = 'Please upload a certificate file.';
			return;
		}

		submitting = true;
		errorMsg = '';

		// Simulate API file upload & verification queue submission
		setTimeout(() => {
			submitting = false;
			submitSuccess = true;

			// Clear form inputs
			setTimeout(() => {
				activityName = '';
				activityCategory = '';
				activityDate = '';
				organizerName = '';
				eventLevel = '';
				certificateNumber = '';
				issueDate = '';
				participationType = '';
				achievementDescription = '';
				selectedFile = null;
				submitSuccess = false;
			}, 3000);
		}, 2000);
	}

	function handleCancel() {
		// Reset states and clear form
		activityName = '';
		activityCategory = '';
		activityDate = '';
		organizerName = '';
		eventLevel = '';
		certificateNumber = '';
		issueDate = '';
		participationType = '';
		achievementDescription = '';
		selectedFile = null;
		errorMsg = '';
	}
</script>

<div class="space-y-6 w-full animate-fade-in font-sans">
	{#if submitSuccess}
		<div
			transition:fade={{ duration: 300 }}
			class="bg-emerald-50 border border-emerald-250 p-6 rounded-xl text-center space-y-4 max-w-xl mx-auto my-12 shadow-xs"
		>
			<div
				class="w-14 h-14 bg-emerald-100 text-emerald-800 flex items-center justify-center rounded-full mx-auto border-2 border-emerald-200 animate-bounce"
			>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					stroke-width="3"
					class="w-7 h-7"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
				</svg>
			</div>
			<div>
				<h3 class="text-lg font-bold text-emerald-900 font-serif">
					Certificate Submitted Successfully!
				</h3>
				<p class="text-xs text-emerald-650 mt-1 max-w-sm mx-auto leading-relaxed">
					Your certificate is now in the verification audit queue. The iSPARC coordinator will
					review and verify your credits shortly.
				</p>
			</div>
		</div>
	{/if}

	<div class="grid grid-cols-1 lg:grid-cols-12 gap-6 items-start">
		<!-- Left Column: Form -->
		<form
			onsubmit={handleSubmit}
			class="lg:col-span-8 bg-white border border-slate-200 rounded-xl overflow-hidden shadow-xs"
		>
			<div class="p-5 border-b border-slate-100">
				<h2 class="text-base font-bold font-serif text-[#0B1535]">Certificate Submission Form</h2>
				<p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest mt-0.5">
					Please provide accurate activity particulars
				</p>
			</div>

			<div class="p-6 space-y-6">
				{#if errorMsg}
					<div
						transition:slide={{ duration: 200 }}
						class="p-4 bg-rose-50 border border-rose-200 text-rose-800 text-xs font-semibold rounded-lg"
					>
						{errorMsg}
					</div>
				{/if}

				<!-- Section: Activity Information -->
				<div class="space-y-4">
					<h3
						class="text-[10px] font-bold text-slate-400 tracking-wider uppercase border-b border-slate-100 pb-2"
					>
						Activity Information
					</h3>

					<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
						<!-- Activity Name -->
						<div class="flex flex-col gap-1.5">
							<label
								for="activity-name"
								class="text-[10px] font-bold text-slate-700 tracking-wider"
							>
								ACTIVITY NAME <span class="text-[#881B1B]">*</span>
							</label>
							<input
								id="activity-name"
								type="text"
								bind:value={activityName}
								placeholder="e.g. National Hackathon 2026"
								class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
								required
							/>
						</div>

						<!-- Activity Category -->
						<div class="flex flex-col gap-1.5">
							<label
								for="activity-category"
								class="text-[10px] font-bold text-slate-700 tracking-wider"
							>
								ACTIVITY CATEGORY
							</label>
							<input
								id="activity-category"
								type="text"
								bind:value={activityCategory}
								placeholder="Auto-filled after selection"
								class="w-full px-3.5 py-2.5 bg-slate-50/80 rounded-lg border border-slate-200 text-[13px] text-slate-500 font-semibold cursor-not-allowed focus:outline-none"
								readonly
							/>
						</div>

						<!-- Activity Date -->
						<div class="flex flex-col gap-1.5">
							<label
								for="activity-date"
								class="text-[10px] font-bold text-slate-700 tracking-wider"
							>
								ACTIVITY DATE <span class="text-[#881B1B]">*</span>
							</label>
							<input
								id="activity-date"
								type="date"
								bind:value={activityDate}
								class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
								required
							/>
						</div>

						<!-- Organizer Name -->
						<div class="flex flex-col gap-1.5">
							<label
								for="organizer-name"
								class="text-[10px] font-bold text-slate-700 tracking-wider"
							>
								ORGANIZER NAME <span class="text-[#881B1B]">*</span>
							</label>
							<input
								id="organizer-name"
								type="text"
								bind:value={organizerName}
								placeholder="Enter organizing institution"
								class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
								required
							/>
						</div>
					</div>

					<!-- Event Level -->
					<div class="flex flex-col gap-1.5">
						<label for="event-level" class="text-[10px] font-bold text-slate-700 tracking-wider">
							EVENT LEVEL <span class="text-[#881B1B]">*</span>
						</label>
						<select
							id="event-level"
							bind:value={eventLevel}
							class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
							required
						>
							<option value="" disabled selected hidden>Select Event Level</option>
							{#each eventLevels as level}
								<option value={level}>{level}</option>
							{/each}
						</select>
					</div>
				</div>

				<!-- Section: Upload File -->
				<div class="space-y-4 pt-2">
					<h3
						class="text-[10px] font-bold text-slate-400 tracking-wider uppercase border-b border-slate-100 pb-2"
					>
						Upload Certificate
					</h3>

					<div
						class="border-2 border-dashed rounded-xl p-8 text-center transition-all duration-200 relative flex flex-col items-center justify-center bg-slate-50/40
							{isDragOver ? 'border-[#881B1B] bg-[#881B1B]/5' : 'border-slate-200 hover:border-slate-350'}
						"
						ondragover={handleDragOver}
						ondragleave={handleDragLeave}
						ondrop={handleDrop}
						role="presentation"
					>
						<input
							id="certificate-upload-file"
							type="file"
							accept=".pdf,.jpg,.jpeg,.png"
							onchange={handleFileChange}
							class="absolute inset-0 w-full h-full opacity-0 cursor-pointer z-10"
						/>

						<div class="w-12 h-12 rounded-full bg-slate-100 flex items-center justify-center mb-3">
							<svg
								class="w-6 h-6 text-slate-450"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								stroke-width="2.2"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
								/>
							</svg>
						</div>

						{#if selectedFile}
							<div class="z-20 relative space-y-1">
								<p class="text-xs font-bold text-[#881B1B]">{selectedFile.name}</p>
								<p class="text-[10px] text-slate-405">
									{(selectedFile.size / 1024 / 1024).toFixed(2)} MB
								</p>
								<button
									type="button"
									onclick={removeSelectedFile}
									class="mt-2.5 px-3 py-1 bg-rose-50 border border-rose-200 text-rose-700 hover:bg-rose-100 rounded-md text-[10px] font-bold uppercase transition duration-150 cursor-pointer"
								>
									Remove File
								</button>
							</div>
						{:else}
							<p class="text-xs font-bold text-slate-800">Drag & Drop your certificate here</p>
							<p class="text-[10px] text-slate-400 mt-0.5">or</p>

							<div class="mt-3.5 z-20 relative">
								<span
									class="inline-flex items-center justify-center bg-[#0B1535] hover:bg-[#0B1535]/90 text-white font-bold text-xs tracking-wider uppercase px-5 py-2.5 rounded-lg shadow-sm transition duration-150"
								>
									Browse Files
								</span>
							</div>

							<p class="text-[10px] text-slate-400 mt-3">
								Supported: PDF, PNG, JPEG, JPG - Max size: 10 MB
							</p>
						{/if}
					</div>
				</div>

				<!-- Section: Certificate Information -->
				<div class="space-y-4 pt-2">
					<h3
						class="text-[10px] font-bold text-slate-400 tracking-wider uppercase border-b border-slate-100 pb-2"
					>
						Certificate Information
					</h3>

					<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
						<!-- Certificate Number -->
						<div class="flex flex-col gap-1.5">
							<label
								for="certificate-number"
								class="text-[10px] font-bold text-slate-700 tracking-wider"
							>
								CERTIFICATE NUMBER <span class="text-[#881B1B]">*</span>
							</label>
							<input
								id="certificate-number"
								type="text"
								bind:value={certificateNumber}
								placeholder="e.g. CERT-2025-001"
								class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
								required
							/>
						</div>

						<!-- Issue Date -->
						<div class="flex flex-col gap-1.5">
							<label for="issue-date" class="text-[10px] font-bold text-slate-700 tracking-wider">
								ISSUE DATE <span class="text-[#881B1B]">*</span>
							</label>
							<input
								id="issue-date"
								type="date"
								bind:value={issueDate}
								class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
								required
							/>
						</div>
					</div>

					<!-- Participation Type -->
					<div class="flex flex-col gap-1.5">
						<label
							for="participation-type"
							class="text-[10px] font-bold text-slate-700 tracking-wider"
						>
							PARTICIPATION TYPE <span class="text-[#881B1B]">*</span>
						</label>
						<select
							id="participation-type"
							bind:value={participationType}
							class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
							required
						>
							<option value="" disabled selected hidden>Select Participation Type</option>
							{#each participationTypes as type}
								<option value={type}>{type}</option>
							{/each}
						</select>
					</div>

					<!-- Achievement Description -->
					<div class="flex flex-col gap-1.5">
						<label
							for="achievement-desc"
							class="text-[10px] font-bold text-slate-700 tracking-wider"
						>
							ACHIEVEMENT DESCRIPTION <span class="text-[#881B1B]">*</span>
						</label>
						<textarea
							id="achievement-desc"
							rows="4"
							bind:value={achievementDescription}
							placeholder="Provide a brief description of your participation or achievement."
							class="w-full px-3.5 py-2.5 bg-white rounded-lg border border-slate-250 text-[13px] text-slate-800 placeholder:text-slate-400 focus:outline-none focus:border-[#881B1B] focus:ring-2 focus:ring-[#881B1B]/10 transition-all duration-200"
							required></textarea>
					</div>
				</div>
			</div>

			<!-- Form Actions -->
			<div
				class="p-6 border-t border-slate-100 bg-slate-50/50 flex items-center justify-end gap-3.5"
			>
				<button
					type="button"
					onclick={handleCancel}
					class="px-5 py-2.5 border border-transparent text-[#6B7280] hover:text-[#111827] font-semibold text-xs tracking-wider uppercase rounded-lg transition duration-150 cursor-pointer"
				>
					Cancel
				</button>

				<button
					type="button"
					onclick={() => alert('Draft Saved Successfully!')}
					class="px-5 py-2.5 border border-slate-250 hover:border-slate-800 text-slate-705 hover:bg-slate-50 font-bold text-xs tracking-wider uppercase rounded-lg transition duration-150 cursor-pointer"
				>
					Save Draft
				</button>

				<button
					type="submit"
					disabled={submitting}
					class="px-6 py-2.5 bg-[#0B1535] hover:bg-[#0B1535]/90 disabled:bg-slate-200 disabled:text-slate-400 text-white font-bold text-xs tracking-wider uppercase rounded-lg transition duration-150 flex items-center gap-2 cursor-pointer"
				>
					{#if submitting}
						<svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
							<circle
								class="opacity-25"
								cx="12"
								cy="12"
								r="10"
								stroke="currentColor"
								stroke-width="4"
							></circle>
							<path
								class="opacity-75"
								fill="currentColor"
								d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
							></path>
						</svg>
						SUBMITTING...
					{:else}
						SUBMIT CERTIFICATE
					{/if}
				</button>
			</div>
		</form>

		<!-- Right Column: Guidelines -->
		<div class="lg:col-span-4 space-y-6">
			<!-- Submission Guidelines Card -->
			<div class="bg-white border border-slate-200 p-6 rounded-xl shadow-xs">
				<div class="border-b border-slate-100 pb-4 mb-5">
					<h3 class="text-base font-bold font-serif text-[#0B1535]">Submission Guidelines</h3>
					<p class="text-[10px] font-semibold text-slate-400 uppercase tracking-widest mt-0.5">
						extracurricular audit rules
					</p>
				</div>

				<ul class="space-y-4">
					{#each guidelines as item}
						<li class="flex items-start gap-3">
							<div
								class="w-5 h-5 rounded-full bg-emerald-50 border border-emerald-250 flex items-center justify-center shrink-0 text-emerald-600"
							>
								<svg
									class="w-3.5 h-3.5"
									fill="none"
									stroke="currentColor"
									stroke-width="3.5"
									viewBox="0 0 24 24"
								>
									<polyline points="20 6 9 17 4 12" />
								</svg>
							</div>
							<span class="text-xs text-slate-655 font-medium leading-normal">{item}</span>
						</li>
					{/each}
				</ul>
			</div>

			<!-- Academic Integrity Notice -->
			<div
				class="bg-white border-l-4 border-[#881B1B] border border-y-slate-200 border-r-slate-200 p-6 rounded-r-xl shadow-xs space-y-3"
			>
				<h4 class="text-xs font-bold text-[#881B1B] uppercase tracking-wider">
					ACADEMIC INTEGRITY NOTICE
				</h4>
				<p class="text-xs text-slate-600 leading-relaxed">
					Submission of forged or altered certificates may result in disciplinary action and
					permanent rejection of extracurricular credits.
				</p>
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
