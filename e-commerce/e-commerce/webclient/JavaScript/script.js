var allProductsGlobal;
var categoryVal;
var category;

const data = new Request("http://localhost:8000/products");

async function getData() {
  let response = await fetch(data);
  let allProducts = await response.json();
  // make allProducts variable global
  allProductsGlobal = allProducts;
  return allProductsGlobal;
};

function buildHtml() {
  if (categoryVal === "all products") {
    let html="";
    for (i = 0; i < allProductsGlobal.length; i++) { 
      html+= `
      <div id=${allProductsGlobal[i].product_id} class="interior div">
          <a href=${allProductsGlobal[i].info_link}>
              <img src=${allProductsGlobal[i].img_src}>
          </a>
          <h4 id=productname>${allProductsGlobal[i].product_name}</h4>
          <p id=productdesc>${allProductsGlobal[i].product_desc}</p><br>
          <p id=price>$${allProductsGlobal[i].product_price}</p> 
      </div>`
  }
    document.getElementById("products").innerHTML=html;
  } else {
    let html="";
      for (i = 0; i < allProductsGlobal.length; i++) { 
        if (allProductsGlobal[i].categories != categoryVal) {continue;}
        html+= `
        <div id=${allProductsGlobal[i].product_id} class="interior div">
            <a href=${allProductsGlobal[i].info_link}>
                <img src=${allProductsGlobal[i].img_src}>
            </a>
            <h4 id=productname>${allProductsGlobal[i].product_name}</h4>
            <p id=productdesc>${allProductsGlobal[i].product_desc}</p><br>
            <p id=price>$${allProductsGlobal[i].product_price}</p> 
        </div>`
    }
      document.getElementById("products").innerHTML=html;
  }
  
};

async function catValues() {
  await getData();
  category = document.getElementById("shop");
  categoryVal = category.options[category.selectedIndex].value;
  // return categoryVal;
  console.log(categoryVal);
  buildHtml();
};

// function filterProducts() {
//     // category = document.getElementById("shop");
//     // categoryVal = category.options[category.selectedIndex].value;
//     if (categoryVal === "all products") {
//       console.log("all products");
//       buildHtml();
//     } else if (categoryVal === "brewers") {
//       buildHtml();
//     } else if (categoryVal === "ceramics") {
//       buildHtml();
//     } else if (categoryVal === "tools") {
//       buildHtml();
//       console.log("tools");
//     } else if (categoryVal === "accessories") {
//       buildHtml();
//     } 
// };
