<script setup>
import { buildPath } from '../api.js'
import plus from '../assets/img/cart/plus.svg'
import minus from '../assets/img/cart/minus.svg'
</script>

<script>
import { mapGetters, mapActions } from 'vuex';
import { buildPath } from '../api.js'

export default {
  computed: {
    ...mapGetters([
      'getProductsInCart',
    ]),
  },

  methods: {
    ...mapActions([
        'addProduct',
        'removeProduct',
    ]),
    hasProduct() {
      return this.getProductsInCart.size > 0;
    },
    addProductToCart(product) {
      this.addProduct(product);
    },
    removeProductFromCart(product) {
      this.removeProduct(product);
    },
    totalPrice() {
        let current = 0;
        function logMapElements(value, key, map) {
            current += (value.product.discount_price ? value.product.discount_price : value.product.price) * value.count
        }
        this.getProductsInCart.forEach(logMapElements);
        return current
    },
  },
};
</script>

<template>
    <div class="container shopping-cart">
        <div class="container-flex-row">
            <div class="cart-items container-flex-column">
                <h1>Shopping Cart</h1>
                <div v-for="[index, product] in getProductsInCart" class="cart-item container-flex-row">
                    <router-link :to="`/api/product/${product.product.id}`">
                    <img style="height: 128px" :src="(product.product.image)">
                    </router-link>
                    <div class="item-title">{{product.product.model}}</div>
                    <div class="item-counter container-flex-row">
                        <button class="count-down" @click.native="removeProductFromCart(product.product)"><img :src="minus"></button>
                        <input class="count-input" type="tel" :value="product.count"></input>
                        <button class="count-up" @click.native="addProductToCart(product.product)"><img :src="plus"></button>
                    </div>
                    <div class="item-price container-flex-row"><p>${{product.product.discount_price ? product.product.discount_price : product.product.price}}</p></div>
                </div>
            </div>
            <div class="cart-summary container-flex-column">
                <h2>Order Summary</h2>
                <div class="cart-total container-flex-row"><p style="padding-right: 20%; padding-top: 1rem;">Total</p><p>${{totalPrice()}}</p></div>
                <button class="cart-checkout">Checkout</button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.shopping-cart {
    padding-top: 4%;
    padding-bottom: 4%;
}

.cart-items {
    width: 50%;
}
.cart-items > h1 {
    font-family: "Abel", sans-serif;
    font-size: 24px;
    font-weight: 400;
    line-height: 24px;
    text-align: left;
    margin-bottom: 10%;
}
.cart-item {
    width: 100%;
}
.item-title {
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: left;;
    vertical-align: middle;
    padding-top: 3.8rem;
    padding-right: 4%;
}
.count-input {
    border: none;
    background: transparent;
    outline: none;
    margin: 0 auto;
    text-align: center;
    width: fit-content;
    max-width: 64px;
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 16px;
    text-align: center;
}
.item-price {
    font-family: "ABeeZee", sans-serif;
    font-size: 20px;
    font-style: italic;
    font-weight: 400;
    line-height: 32px;
    letter-spacing: 0.03em;
    text-align: left;
    vertical-align: middle;
    padding-top: 1rem;
    padding-right: 2rem;
}

.cart-summary {
    padding-top: 8%;
    padding-bottom: 8%;
    padding-left: 8%;
    padding-right: 8%;
    width: 50%;
}
.cart-summary > h2 {
    font-family: "ABeeZee", sans-serif;
    font-size: 20px;
    font-style: italic;
    font-weight: 400;
    line-height: 16px;
    text-align: left;
}
.cart-summary > div {
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    letter-spacing: 0.03em;
    text-align: left;
}
.cart-total {
    padding-top: 8%;
    width: 100%;
    text-align: left;
}
.cart-checkout {
    padding: 16px 56px 16px 56px;
    gap: 8px;
    border-radius: 5%;
    opacity: 0px;
    margin-top: 8%;
    background: #000000;
    color: #FFFFFF;
    font-family: "ABeeZee", sans-serif;
    font-size: 16px;
    font-style: italic;
    font-weight: 400;
    line-height: 24px;
    text-align: center;
}
</style>
