<template>
    <!-- footer section start -->
    <div class="container-footer">
        <div class="container container-flex-column">
            <div class="footer-info container-flex-row">
                <div class="footer-cyber container-flex-row">
                    <img src="../assets/img/footer/logo_vector.svg" style="padding-top: 0%; height: 15%">
                    <p class="footer-cyber-text">We are a residential interior design firm located in Portland. Our boutique-studio offers more than</p>
                </div>
                <div class="footer-services">
                    <h2 style="font-size: 16px;">Services</h2>
                    <div class="footer-services-links">
                        <li><a href="#">Bonus program</a></li>
                        <li><a href="#">Gift cards</a></li>
                        <li><a href="#">Credit and payment</a></li>
                        <li><a href="#">Service contracts</a></li>
                        <li><a href="#">Non-cash account</a></li>
                        <li><a href="#">Payment</a></li>
                    </div>
                </div>
                <div class="footer-assistance">
                    <h2 style="font-size: 16px;">Assistance to the buyer</h2>
                    <div class="footer-assistance-links">
                        <li><a href="#">Find an order</a></li>
                        <li><a href="#">Terms of delivery</a></li>
                        <li><a href="#">Exchange and return of goods</a></li>
                        <li><a href="#">Guarantee</a></li>
                        <li><a href="#">Frequently asked questions</a></li>
                        <li><a href="#">Terms of use of the site</a></li>
                    </div>
                </div>
            </div>
            <div class="footer-links">
                <a href="#"><img class="footer-links-icon" src="../assets/img/footer/twitter.svg"></a>
                <a href="#"><img class="footer-links-icon" src="../assets/img/footer/facebook.svg"></a>
                <a href="#"><img class="footer-links-icon" src="../assets/img/footer/tiktok.svg"></a>
                <a href="#"><img class="footer-links-icon" src="../assets/img/footer/instagram.svg"></a>
            </div>
    
        </div>
    </div>
    <!-- footer section end -->
</template>

<style scoped>
.container-footer {
    width: 100%;
    height: fit-content;
    background-color: black;
    color: #CFCFCF;
    padding-top: 2%;
}

.footer-info {
    align-self: flex-start;
    width: 100%;
}

.footer-cyber {
    width: 30%;
    vertical-align: top;
    padding-left: 0;
}
.footer-cyber-text {
    font-family: "ABeeZee", sans-serif;
    font-size: 14px;
    font-style: italic;
    font-weight: 400;
    line-height: 23.87px;
    text-align: left;
}

.footer-services {
    width: 30%;
    font-family: "Abel", sans-serif;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 32px;
    text-align: left;
    align-self: flex-end;
}
.footer-assistance {
    width: 30%;
    font-family: "Abel", sans-serif;
    font-size: 14px;
    font-style: normal;
    font-weight: 400;
    line-height: 32px;
    text-align: left;
    align-self: flex-end;
}

ul, li {
  list-style: none;
}

.footer-links-icon {
    padding-right: 2%;
    padding-top: 0%;
}
</style>
