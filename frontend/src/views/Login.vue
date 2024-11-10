<template>
    <div class="login container container-flex-column">
        <form @submit.prevent="authenticate()" class="login-form container-flex-column">
            <img src="https://cdn-icons-png.flaticon.com/512/310/310818.png" alt="Login" class="login-form-item container-flex-row" style="width: 12%;">
            <div class="login-form-item container-flex-row">
                <label for="">E-mail</label>
                <input v-model="auth.email" type="email">
            </div>
    
            <div class="login-form-item container-flex-row">
                <label for="">Password</label>
                <input v-model="auth.password" type="password">
            </div>
    
            <div class="login-form-item container-flex-row">
                <button type="submit">Login</button>
            </div>
        </form>
    </div>
    </template>
    
    <script>
    import { userLogin } from '../api.js'
    
    export default {
        data () {
            return {
                auth: { email: '', password: '' },
                user: {},
            }
        },
        methods: {
            /**
             * Attempts to authenticate the user
             *
             * @return {mixed}
             */
            async authenticate () {
                let credentials = this.auth
                console.log(credentials)
    
                const data = await userLogin(credentials)
                console.log(data)
    
                if (data.statusCode === 0) {
                    window.localStorage.setItem('accessToken', data.accessToken)
                    window.localStorage.setItem('authUser', JSON.stringify(data.displayName))
                    
                     this.$router.push({
                        path: `/`
                    })
                } else {
                    alert("Invalid Email or Password")
                }
            }
        }
    }
    </script>
    
    <style scoped>
    .login-form {
        width: 100%;
    }
    .login-form-item{
        width: 30%;
        margin-top: 1rem;
        margin-bottom: 1rem;
        cursor: default;
    }
    .login-form-item label {
        font-family: "ABeeZee", sans-serif;
        font-weight: 400;
        font-style: bold;
        width: 25%;
    }
    .login-form-item input {
        border: none;
        outline: none;
        width: 70%;
        height: 2rem;
        border-radius: 5%;
        opacity: 0px;
        background: #F5F5F5;
        font-family: "ABeeZee", sans-serif;
        font-weight: 400;
        font-style: italic;
        font-size: 14px;
        padding-left: 1rem;
        padding-right: 1rem;
    }
    .login-form-item button {
        padding-top: 1rem;
        padding-bottom: 1rem;
        width: 100%;
        border-radius: 2%;
        font-family: "ABeeZee", sans-serif;
        font-weight: 400;
        font-style: bold;
        font-size: 32px;
        text-align: center;
    }
    .login-form-item button:hover {
        background-color:black;
        color: white;
        cursor: pointer;
    }
    </style>