<script setup>
import Switch from './Switch.vue'
import Serch from './Serch.vue'
import ImgCard from './ImgCard.vue'
import axios from 'axios'
import { onMounted, ref } from 'vue'

const list = ref([])

onMounted(async () => {
	try {
		list.value = await (await axios.get("./load")).data;
	} catch {
		console.log("loading failed")
	}
})
</script>

<template lang="pug">
header.pad
	.logo tetradka
	Switch.switch
	Serch.serch

.page.pad
	ImgCard(v-for="img in list" :id="img.id" :src="img.src")
</template>

<style scoped>

.switch {
	margin: 0 0.8rem;
}

header {
	display: flex;
	flex-flow: row nowrap;
	margin-top: 1vh;
	margin-bottom: 1.5rem;
	font-size: 1rem;
}

.page {
	/* background: snow; */
	/* height: 100vh; */
	display: flex;
	flex-flow: row wrap;
	justify-content: flex-start;
	align-items: flex-start;
	align-content: flex-start;
	gap: 1.5rem 2rem;
}

.pad {
	margin-left: auto;
	margin-right: auto;
	width: 80%;
}

.logo {
	font-family: 'PT Sans';
	font-size: 3rem;
	font-weight: 700;
	font-style: italic;
	color: var(--brand);
	user-select: none;
	width: min-content;
}

</style>

