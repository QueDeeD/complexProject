<script setup>

import { ref } from 'vue'
import { getCategories } from '../api.js'

const loading = ref(false)
const post = ref(null)
const error = ref(null)

async function fetchData() {
    error.value = post.value = null
    loading.value = true
  
    try {
        post.value = await getCategories()
    } catch (err) {
        error.value = err.toString()
    } finally {
        loading.value = false
    }
}

fetchData()
</script>

<template>
<div class="categories container container-flex-column">
    <div class="categories-nav container-flex-row">
        <h2>Browse By Category</h2>
        <div class="nav-arrows">
            <img src="../assets/img/category/arrow_left.svg" style="padding-left: 32px;">
            <img src="../assets/img/category/arrow_right.svg" style="padding-left: 32px;">
        </div>
    </div>
    <div v-if="post" class="categories-cnt container-flex-row">
        <div v-for="category in post" class="category-card">
            <router-link :to="`/api/category/${category.id}`"><img :src="category.image" style="width: 50%;"><p>{{ category.displayName }}</p></router-link>
        </div>
        <router-view></router-view>
    </div>
</div>
</template>

<style scoped>
.categories {
    height: fit-content;
    background: #FAFAFA;
    align-items: center;
    align-content: center;
    padding-left: 5%;
    padding-right: 5%;
}
.categories-nav {
    width: 100%;
    font-family: "ABeeZee", sans-serif;
    font-size: 24px;
    font-style: italic;
    font-weight: 400;
    line-height: 32px;
    letter-spacing: 0.01em;
    text-align: left;
    height: fit-content;
    margin-top: 32px;
    padding-bottom: 2%;
}

.categories-cnt {
    width: 100%;
    height: fit-content;
    margin-bottom: 32px;
}
.category-card {
    background: #EDEDED;
    width: 160px;
    height: 128px;
    padding-top: 1%;
    gap: 8px;
    border-radius: 5%;
    opacity: 0px;
    border-radius: 5%;
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: center;
    align-items: center;
    margin-top: 8px;
    margin-bottom: 8px;
}
</style>
