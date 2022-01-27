<script setup>
import { onMounted } from 'vue'
const props = defineProps({
	id: Number,
	src: String,
});
onMounted(()=>{
	console.log(props)
})

</script>

<template lang="pug">
.frame 
	img(:id="props.id" :src="props.src")

</template>

<style scoped>
.frame {
	box-shadow:
		0px 1.9px 1.9px rgba(0, 0, 0, 0.035),
		0px 15px 15px rgba(0, 0, 0, 0.07)
	;
	background: black;
	border-top: 1px solid #fffffa;
	width: 200px;
	min-height: 10rem;
}

.frame img {
	user-select: none;
	height: 100%;
	width: 100%;
}

</style>