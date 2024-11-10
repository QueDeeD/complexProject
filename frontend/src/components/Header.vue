<script setup>

</script>

<template>
    <div class="header-top container container-flex-row">
        <div class="header-logo">
            <router-link to="/"><img src="../assets/img/header/logo_vector.svg"></router-link>
        </div>
        <div class="header-search container-flex-row">
            <div class="header-image"><img src="../assets/img/header/search_vector.svg" ></div>
            <input type="search" class="header-input" title="Search" value="Search">
        </div>
        <div class="header-icons container-flex-row">
            <div class="header-icon">
                <img src="../assets/img/header/favorites_vector.svg" width="30px">
            </div>
            <div class="header-icon">
                <router-link to="/cart"><img src="../assets/img/header/cart_vector.svg " width="30px"></router-link>
            </div>
            <div class="header-icon">
                <router-link to="/login"><img src="../assets/img/header/login_vector.jpg" width="30px"></router-link>
            </div>
        </div>
    </div>
</template>

<style scoped>
.header-top {
    margin-bottom: 2%;
    margin-top: 2%;
}
.header-logo {
    align-self: center;
}
.header-search {
    width: 30%;
    height: fit-content;
    border-radius: 5%;
    opacity: 0px;
    background: #F5F5F5;
    font-family: "ABeeZee", sans-serif;
    font-weight: 400;
    font-style: italic;
    font-size: 14px;
}
.header-input {
    border: none;
    background: transparent;
    outline: none;
    margin: 0 auto;
    text-align: left;
    padding-right: 2%;
    padding-left: 2%;
    width: 60%;
    padding-top: 5%;
    padding-bottom: 5%;
}
.header-image {
    width: 20%;
    padding-left: 5%;
    align-self: center;
}
.header-icon {
    padding-left: 1rem;
    align-self: center;    
}
</style>
