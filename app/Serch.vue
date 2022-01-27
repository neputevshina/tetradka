<script setup>
import { ref } from 'vue'
const fuzzy = ref(true)
</script>

<template lang="pug">
div
	.toggle(@click="fuzzy=!fuzzy" :class="{yes:fuzzy}" title="Нечёткий поиск")
		span
	input(type=search, placeholder="Поиск...", height=0, width=0)
</template>

<style scoped>
div {
	width: 100%;
	display: flex;
	height: min-content;
}

.toggle {
	margin: auto 0;
	width: auto;
}

.toggle.yes span {
	background: var(--glittery);
	transition: ease-in 30ms;
}

.toggle span {
	margin: auto auto;
	mask: url("./assets/fuzzy.svg");
	width: 16px;
	height: 13px;
	background: var(--greylike);
	transition: ease-out 30ms;
}

div input {
	padding: 0.2rem 0;
	margin-left: 0.2rem;
	height: 1rem;
	width: 100%;
}

div input::placeholder {
	color: var(--greylike);
	opacity: 1; /* because fucking vue interprets 100% as 1% in release mode */
}
</style>
