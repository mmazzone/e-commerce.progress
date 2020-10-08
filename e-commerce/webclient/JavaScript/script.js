// declare global variables
var allProductsGlobal;
var categoryVal;
var category;

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
then stores category values to use in the buildHtml function */
async function catValues() {
  await getData();
  category = document.getElementById("shop");
  categoryVal = category.options[category.selectedIndex].value;
  buildHtml();
}

// funtion iterates through each item then displays it into othe HTML
function buildHtml() {
  let html = "";
  if (categoryVal === "all products") {
    for (i = 0; i < allProductsGlobal.length; i++) {
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
    // populates html
    document.getElementById("products").innerHTML = html;

    // function compares the category value allowing products to filter by their category and then displays those into the HTML
  } else {
    let html = "";
    for (i = 0; i < allProductsGlobal.length; i++) {
      /* if the product category is different than the category selected it will continue and iterate through the next value until 
        category is matched */
      if (allProductsGlobal[i].categories != categoryVal) {
        continue;
      }
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
    // populates html
    document.getElementById("products").innerHTML = html;
  }
}
