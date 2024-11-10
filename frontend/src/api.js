const endpoint = 'http://localhost:8090'
const target_userlogin = '/user/login'
const target_categories = '/categories' //'/api/category/'
const target_category = '/category/list/' //'/api/category/'
//const target_subcategory = '/api/subcategory/'
const target_products = '/product/get/' //'/api/products/'

function buildPath(target) {
    let ep = endpoint
    return ep.concat('/', target)
}

async function userLogin(validation) {
    let ep = endpoint
    const remote = ep.concat('', `${target_userlogin}?username=${validation.email}&password=${validation.password}`)
    console.log(remote)
    const requestOptions = {
        method: "POST",
    };
    const response = await fetch(remote, requestOptions)
    return response.json();
}

async function getCategories() {
    let ep = endpoint
    const remote = ep.concat('', target_categories)
    console.log(remote)
    const response = await fetch(remote)
    return response.json();
}

async function getCategory(id) {
    let ep = endpoint
    const remote = ep.concat('', `${target_category}${id}`)
    console.log(remote)
    const response = await fetch(remote)
    return response.json();
}

async function getProduct(id) {
    let ep = endpoint
    const response = await fetch(ep.concat('', `${target_products}${id}`))
    return response.json();
}

export {buildPath, userLogin, getCategories, getCategory, getProduct};
