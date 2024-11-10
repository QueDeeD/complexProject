<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getProduct, buildPath } from '../api.js'
import screensize from '../assets/img/product/info/specs/screensize.svg'
import cpu from '../assets/img/product/info/specs/cpu.svg'
import cores from '../assets/img/product/info/specs/cores.svg'
import battery from '../assets/img/product/info/specs/battery.svg'
import delivery from '../assets/img/product/info/misc/delivery.svg'
import stock from '../assets/img/product/info/misc/stock.svg'
import guarantee from '../assets/img/product/info/misc/guarantee.svg'

const route = useRoute()

const loading = ref(false)
const post = ref(null)
const error = ref(null)

const mapIconsSpec = new Map([
    ["Диагональ", screensize],
    ["Разрешение", screensize],
    ["Процессор", cpu],
    ["Аккумулятор", battery],
    ["Питание", battery],
    ["Кол-во ядер", cores]
]);

const getIcon = (elem) => {
    let res = mapIconsSpec.get(elem.characteristic)
    console.log(`${elem.characteristic} : ${res}`)
    return res == undefined ? '' : res
}

watch(() => route.params.id, fetchData, { immediate: true })

async function fetchData(id) {
    error.value = post.value = null
    loading.value = true

    try {
        post.value = await getProduct(id)
    } catch (err) {
        error.value = err.toString()
    } finally {
        loading.value = false
    }
}
</script>

<script>
import { mapActions } from 'vuex';

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
    <div style="background: #FAFAFA;">
    <div v-if="post" class="container container-flex-column">
        <div class="product-intro container-flex-row">
            <div class="product-images">
                <img :src="(post.image)">
            </div>
            <div class="product-info container-flex-column">
                <div class="product-title">{{post.model}}</div>
                <div class="product-price container-flex-row"><p>${{post.discount_price ? post.discount_price : post.price}}</p><s>{{post.discount_price ? `$${post.price}` : ''}}</s></div>
                <div class="product-specs container-flex-row">
                    <div v-for="characteristic in post.characteristics" class="product-spec container-flex-row">
                        <div class="ps-image" ><img :src="getIcon(characteristic)"></div>
                        <div class="ps-text" >
                            <h3>{{characteristic.characteristic}}</h3>
                            <p>{{characteristic.value}}</p>
                        </div>
                    </div>
                </div>
                <div class="product-descr"></div>
                <div class="product-actions container-flex-row">
                    <button class="action-wishlist">Add to Wishlist</button>
                    <button class="action-cart" @click.native="addProductToCart(post)">Add to Cart</button>
                </div>
                <div class="product-misc container-flex-row">
                    <div class="misc-delivery container-flex-row">
                        <div class="ms-image" ><img :src="delivery"></div>
                        <div class="ms-text container-flex-column" >
                            <h3>Free Delivery</h3>
                            <p>1-2 days</p>
                        </div>
                    </div>
                    <div class="misc-in-stock container-flex-row">
                        <div class="ms-image" ><img :src="stock"></div>
                        <div class="ms-text container-flex-column" >
                            <h3>In Stock</h3>
                            <p>{{post.is_available ? 'Today' : 'None'}}</p>
                        </div>
                    </div>
                    <div class="misc-guaranteed container-flex-row">
                        <div class="ms-image" ><img :src="guarantee"></div>
                        <div class="ms-text container-flex-column" >
                            <h3>Guaranteed</h3>
                            <p>1 year</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="product-details container-flex-column">
            <h2>Details</h2>
            <ul v-for="characteristic in post.characteristics" >
                <li><div class="details-flex-row container-flex-row"><p>{{characteristic.characteristic}}</p><p>{{characteristic.value}}</p></div></li>
            </ul>
        </div>
        <div class="product-mark container-flex-column">
            <div class="product-mark-box">
                <h3>{{post.rating}}</h3>
                <p>{{`of ${post.count_review} reviews`}}</p>
                <span :class="post.rating >= 1 ? 'fa fa-star checked' : 'fa fa-star'"></span>
                <span :class="post.rating >= 2 ? 'fa fa-star checked' : 'fa fa-star'"></span>
                <span :class="post.rating >= 3 ? 'fa fa-star checked' : 'fa fa-star'"></span>
                <span :class="post.rating >= 4 ? 'fa fa-star checked' : 'fa fa-star'"></span>
                <span :class="post.rating >= 5 ? 'fa fa-star checked' : 'fa fa-star'"></span>
            </div>
        </div>
    </div>
    </div>
</template>

<style scoped>
.checked {
  color: orange;
}
.product-mark {
    height: 100%;
    width: 80%;
    align-self: center;
    background: #FFFFFF;
    padding-top: 1%;
    padding-bottom: 1%;
    align-content: center;
    align-items: center;
    margin-bottom: 10%;
}
.product-mark-box {
    height: fit-content;
    width: fit-content;
    background: #FAFAFA;
    padding-left: 4%;
    padding-right: 4%;
    padding-top: 4%;
    padding-bottom: 4%;
    align-self: center;
    border-radius: 5%;
}
.product-mark-box > h3 {
    font-family: "ABeeZee", sans-serif;
    font-size: 56px;
    font-style: italic;
    font-weight: 400;
    line-height: 56px;
    text-align: center;
    margin-bottom: 16px;
}
.product-mark-box > p {
    font-family: "ABeeZee", sans-serif;
    font-size: 15px;
    font-style: italic;
    font-weight: 400;
    line-height: 16px;
    text-align: center;
    margin-bottom: 8px;
}
.product-mark-box > span {
    align-self: center;
}

.product-details {
    height: fit-content;
    width: 80%;
    align-self: center;
    margin-top: 64px;
    margin-bottom: 64px;
    background: #FFFFFF;
    padding-left: 4%;
    padding-right: 4%;
    padding-top: 2%;
    padding-bottom: 2%;
}
.product-details > h2 {
    margin-bottom: 64px;
    font-family: "ABeeZee", sans-serif;
    font-size: 24px;
    font-style: italic;
    font-weight: 400;
    line-height: 32px;
    text-align: left;
}
.details-flex-row {
    height: fit-content;
    width: 100%;
    align-self: center;
}
.details-flex-row > p {
    font-family: "Abel", sans-serif;
    font-size: 16px;
    font-weight: 400;
    line-height: 24px;
    text-align: left;
}

.product-intro {
    height: fit-content;
    width: 100%;
    align-self: center;
    align-items: center;
    align-content: center;
    background: #FFFFFF;
}
.product-images {
    width: 40%;
}
.product-info {
    height: fit-content;
    width: 60%;
    align-self: center;
    padding-right: 10%;
}
.product-title {
    font-family: "ABeeZee", sans-serif;
    font-size: 40px;
    font-style: italic;
    font-weight: 400;
    line-height: 40px;
    text-align: left;
}
.product-price {
    width: 100%;
}
.product-price > p {
    font-family: "ABeeZee", sans-serif;
    font-size: 32px;
    font-style: italic;
    font-weight: 400;
    line-height: 48px;
    letter-spacing: 0.03em;
    text-align: left;
}
.product-price > s {
    font-family: "Abel", sans-serif;
    font-size: 24px;
    font-weight: 400;
    line-height: 32px;
    letter-spacing: 0.03em;
    text-align: left;
}

.product-specs {
    height: fit-content;
    width: fit-content;
    align-items: center;
    align-content: center;
}
.product-spec {
    height: fit-content;
    width: fit-content;
    padding: 16px 0px 0px 0px;
    gap: 8px;
    opacity: 0px;
    background: #F4F4F4;
    margin-bottom: 8px;
    margin-top: 8px;
    border-radius: 5%;
    align-items: center;
    align-content: center;
    padding-left: 2%;
    padding-right: 3%;
    padding-bottom: 2%;
    margin-right: 1%;
}
.ps-text > h3 {
    font-family: "Abel", sans-serif;
    font-size: 14px;
    font-weight: 400;
    line-height: 16px;
    text-align: left;
    color: #4E4E4E;
    opacity: 50%;
}
.ps-text > p {
    font-family: "ABeeZee", sans-serif;
    font-size: 14px;
    font-style: italic;
    font-weight: 400;
    line-height: 16px;
    text-align: left;
    color: #4E4E4E;
}
.ps-image {
    width: fit-content;
    height: 100%;
    padding-left: 0;
}
.ps-text {
    width: fit-content;
    height: 100%;
    padding-left: 0;
    padding-bottom: 2%;
}

.product-actions {
    width: 100%;
    height: 56px;
    align-items: center;
    align-content: center;
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: center;
    margin-top: 32px;
    margin-bottom: 32px;
}
.action-wishlist {
    width: 48%;
    height: 100%;
    gap: 8px;
    border-radius: 5%;
    border: 1px 0px 0px 0px;
    opacity: 0px;
    border: 1px solid #000000;
    background: #FFFFFF;
    text-align: center;
    margin-right: 2%;
}
.action-cart {
    width: 48%;
    height: 100%;
    gap: 8px;
    border-radius: 5%;
    border: 1px 0px 0px 0px;
    opacity: 0px;
    border: 1px solid #000000;
    background: #000000;
    color: #FFFFFF;
    text-align: center;
    margin-left: 2%;
}

.product-misc {
    width: 100%;
    height: fit-content;
    align-items: center;
    align-content: center;
    margin-top: 16px;
}
.product-misc > div {
    width: fit-content;
    height: 100%;
    background: #F6F6F6;
    align-items: center;
    align-content: center;
    padding-left: 4%;
    padding-right: 4%;
    padding-top: 3%;
    padding-bottom: 3%;
    margin-top: 1rem;
}
.ms-text > h3 {
    font-family: "ABeeZee", sans-serif;
    font-size: 14px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: left;
    color: #717171;
}
.ms-text > p {
    font-family: "ABeeZee", sans-serif;
    font-size: 14px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: left;
    color: #000000;
}
.ms-image {
    width: fit-content;
    height: 100%;
    padding-left: 0;
    width: 30%;
}
.ms-text {
    width: 70%;
    height: 100%;
    padding-right: 0;
    padding-bottom: 2%;
}
</style>
