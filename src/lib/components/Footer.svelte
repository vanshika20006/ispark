<script lang="ts">
	const year = new Date().getFullYear();
	const currentYear = year < 2026 ? 2026 : year;
</script>

<footer class="py-12 bg-bg-warm border-t border-slate-200">
	<div class="max-w-7xl mx-auto px-6 lg:px-8 text-center space-y-3">
		<p class="text-xs text-slate-500 font-sans leading-relaxed">
			© {currentYear} Institute of Infrastructure & Professional Studies (IIPS), DAVV. All rights reserved.
		</p>
		<p class="text-[10px] font-bold tracking-wider text-slate-400 uppercase font-sans">
			ISPARC — Skill, Personality Advancement & Refinement Cell
		</p>
	</div>
</footer>
