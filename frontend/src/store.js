import { createStore } from 'vuex'

const store = createStore({
    state: {
        cartProducts: new Map(),
        currentProduct: {},
        showModal: false,
        showPopupCart: false,
    },

    getters: {
        getSmartphones: state => state.smartphones,
        getProductsInCart: state => state.cartProducts,
        getCurrentProduct: state => state.currentProduct,
        getShowModal: state => state.showModal,
        getPopupCart: state => state.showPopupCart,
    },

    mutations: {
        ADD_PRODUCT: (state, product) => {
            console.log(`ADD_PRODUCT size ${state.cartProducts.size}`)
            let cached = state.cartProducts.get(product.id)
            if (cached === undefined) {
                const obj = {
                    count: 1,
                    product: product,
                };
                state.cartProducts.set(product.id, obj);
            } else {
                console.log(`ADD_PRODUCT Cached ${product.id} : ${cached.count} items`)
                cached.count += 1;
            }
        },
        REMOVE_PRODUCT: (state, product) => {
            console.log(`REMOVE_PRODUCT size ${state.cartProducts.size}`)
            let cached = state.cartProducts.get(product.id)
            if (cached === undefined) {
            } else {
                console.log(`REMOVE_PRODUCT Cached ${product.id} : ${cached.count} items`)
                cached.count -= 1;
                if (cached.count === 0) {
                    state.cartProducts.delete(product.id)
                }
            }
        },
        CURRENT_PRODUCT: (state, product) => {
            state.currentProduct = product;
        },
        SHOW_MODAL: (state) => {
            state.showModal = !state.showModal;
        },
        SHOW_POPUP_CART: (state) => {
            state.showPopupCart = !state.showPopupCart;
        },
    },

    actions: {
        addProduct: (context, product) => {
            context.commit('ADD_PRODUCT', product);
            console.log(`Product added: ${product.id}`)
        },
        removeProduct: (context, product) => {
            context.commit('REMOVE_PRODUCT', product);
            console.log(`Product removed: ${product.id}`)
        },
        currentProduct: (context, product) => {
            context.commit('CURRENT_PRODUCT', product);
        },
        showOrHiddenModal: (context) => {
            context.commit('SHOW_MODAL');
        },
        showOrHiddenPopupCart: (context) => {
            context.commit('SHOW_POPUP_CART');
        },
    },
});

export { store }
