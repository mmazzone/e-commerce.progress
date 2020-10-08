// declare global variables
var allProductsGlobal;
var priceRange;
var priceRangeValue;

// variable to request data from backend
const data = new Request("http://localhost:8000/products");

// function pulls data - asysnc function will wait to complete until the promise has been fulfilled
async function getData() {
  let response = await fetch(data);
  let allProducts = await response.json();
  // make allProducts variable global
  allProductsGlobal = allProducts;
  return allProductsGlobal;
}
/* function waits for product data to become available, 
then stores priceRange data for use in the searchProducts function. */
async function priceValues() {
  await getData();
  priceRange = document.getElementById("price");
  priceRangeValue = priceRange.options[priceRange.selectedIndex].value;
  searchProducts();
}

// function compares search input to products available to display search results
function searchProducts() {
  var searchInput = document.getElementById("search").value;
  // search input is converted to lowercase
  var inputToLower = searchInput.toLowerCase();
  let html = "";
  for (i = 0; i < allProductsGlobal.length; i++) {
    // product names are also converted to lower case to compare
    if (
      allProductsGlobal[i].product_name.toLowerCase().includes(inputToLower)
    ) {
      if (
        priceRangeValue === "price1" &&
        allProductsGlobal[i].product_price <= 20
      ) {
        html += `
                <div id=${allProductsGlobal[i].product_id} class="interior div">
                    <a href=${allProductsGlobal[i].info_link}>
                        <img src=${allProductsGlobal[i].img_src}>
                    </a>
                    <h4 id=productname>${allProductsGlobal[i].product_name}</h4>
                    <p id=productdesc>${allProductsGlobal[i].product_desc}</p><br>
                    <p id=price>$${allProductsGlobal[i].product_price}</p> 
                </div>`;
      } else if (
        priceRangeValue === "price2" &&
        allProductsGlobal[i].product_price > 20
      ) {
        html += `
                <div id=${allProductsGlobal[i].product_id} class="interior div">
                    <a href=${allProductsGlobal[i].info_link}>
                        <img src=${allProductsGlobal[i].img_src}>
                    </a>
                    <h4 id=productname>${allProductsGlobal[i].product_name}</h4>
                    <p id=productdesc>${allProductsGlobal[i].product_desc}</p><br>
                    <p id=price>$${allProductsGlobal[i].product_price}</p> 
                </div>`;
      }
    }
  }
  // populates html
  document.getElementById("products").innerHTML = html;
}
