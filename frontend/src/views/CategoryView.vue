<script setup>

import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getCategory } from '../api.js'

const route = useRoute()

const loading = ref(false)
const post = ref(null)
const error = ref(null)

watch(() => route.params.id, fetchData, { immediate: true })

async function fetchData(id) {
    error.value = post.value = null
    loading.value = true
  
    try {
        post.value = await getCategory(id)
    } catch (err) {
        error.value = err.toString()
    } finally {
        loading.value = false
    }
}
// https://codesandbox.io/s/route-params-vue-router-examples-mlb14?from-embed=&initialpath=/users/eduardo/posts/1&file=/src/main.js:60-93
// https://router.vuejs.org/guide/essentials/dynamic-matching.html
// https://router.vuejs.org/guide/advanced/data-fetching
// https://codesandbox.io/s/dynamic-component-rendering-using-vuejs-5z6ff?fontsize=14&file=/src/App.vue
</script>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  methods: {
    ...mapActions([
      'addProduct',
    ]),
    addProductToCart(product) {
      this.addProduct(product);
    },
  },

};
</script>

<template>
<div class="category container container-flex-column">
    <div class="category-menu">
    
    </div>
    <div v-if="post" class="category-content container-flex-row">
        <div class="category-content-head container-flex-row">
            <h2 class="category-content-head-title">Selected Products: {{post.length}}</h2>
            <select name="product-opts" id="product-select" class="category-content-head-select">
                <option value="">--Please choose an option--</option>
                <option value="By rating">By rating</option>
                <option value="By price">By price</option>
            </select>
        </div>
        <div class="category-content-body container-flex-row">
            <div v-for="product in post" class="product-card container-flex-column">
                <router-link :to="`/api/product/${product.id}`">
                    <div class="product-card-like">
                        <img src="../assets/img/products/fav.svg">
                    </div>
                    <div class="product-card-image">
                        <img :src="(product.image)" style="width: 50%; ">
                    </div>
                    <p class="product-card-title">{{ product.model }}</p>
                    <div class="product-card-price">${{product.price}}</div>
                </router-link>
                <button class="product-card-action" @click.native="addProductToCart(product)">Buy Now</button>
            </div>
        </div>
        <div class="category-content-pages">
    
        </div>
    </div>
</div>
</template>

<style scoped>
.category-content {
    padding-left: 20%;
    padding-right: 20%;
    margin-bottom: 10%;
}
.category-content-body {
    width: 100%;
}
.product-card {
    background: #F6F6F6;
    width: 266.33px;
    height: 432px;
    padding: 24px 16px 24px 16px;
    gap: 16px;
    border-radius: 9px 0px 0px 0px;
    opacity: 0px;
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: center;
    align-content: center;
    align-items: center;
    margin-top: 16px;
    margin-bottom: 16px;
}
.product-card-like {
    padding-right: 0;
    width: 8%;
}
.product-card-title {
    margin-top: 10%;
    margin-bottom: 10%;
}
.product-card-price {
    font-family: "Abel", sans-serif;
    font-size: 24px;
    font-weight: 400;
    line-height: 24px;
    letter-spacing: 3%;
    letter-spacing: 0.03em;
    text-align: center;
    margin-top: 10%;
}
.product-card-action {
    width: 186px;
    height: 48px;
    padding: 12px 64px 12px 64px;
    gap: 8px;
    border-radius: 8px 0px 0px 0px;
    opacity: 0px;
    background: #000000;
    border-radius: 5%;
    font-family: "ABeeZee", sans-serif;
    font-size: 14px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: center;
    color: #FFFFFF;
}

.category-content-head {
    width: 100%;
    height: 40px;
    gap: 0px;
    opacity: 0px;
    height: fit-content;
}
.category-content-head-title {
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 16px;
    letter-spacing: 0.03em;
    text-align: left;
    margin-top: 16px;
}
.category-content-head-select {
    font-family: "Abel", sans-serif;
    font-size: 15px;
    font-weight: 400;
    line-height: 16px;
    letter-spacing: -0.005em;
    text-align: left;
    padding-right: 0;
    width: 256px;
    height: 40px;
    padding: 8px 16px 8px 16px;
    gap: 0px;
    border-radius: 8px 0px 0px 0px;
    border: 0.5px 0px 0px 0px;
    opacity: 0px;
    border: none;
    background: transparent;
    outline: none;
}
</style>
