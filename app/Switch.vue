<script setup>
import { ref } from 'vue'
const tog = (i) => {
	console.log(selected.value ^ i)
	if ((selected.value ^ (1<<i)) !== 0) {
		selected.value = selected.value ^ (1<<i)
	}
}
const ck = (i) => (selected.value&(1<<i))!==0

const selected = ref(0x01)
</script>

<template lang="pug">
.root
	.button(@click="tog(0)" :class="{active: ck(0)}")
		span Картинки
	.button(:class="{active: ck(1)}")
		span Текст
		.tbd
	.button(:class="{active: ck(2)}")
		span Аудио
		.tbd
	.button(:class="{active: ck(3)}")
		span Видео
		.tbd

</template>

<style scoped>
.root {
	display: flex;
	flex-direction: row;
	width: max-content;
	height: min-content;
}

.button {
	background: inherit;
	padding: 0.2rem 0.4rem;
	height: max-content;
	text-decoration: line-through; 
	color: var(--greylike);
	transition: ease-in 100ms;
	position: relative;
}

.tbd {
	padding: 0.2rem 0.4rem;
	color: red;
	top: 0;
	margin: 0 auto;
	position:absolute;
}

.tbd::after {
	content: "TBD";
}

.button span {
	user-select: none;
}

.button.active {
	text-decoration: none;
	color: unset;
	transition: ease-out 100ms;
}
</style>