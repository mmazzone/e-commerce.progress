




html+=`
<div id="${}" class="interior div1"><!--may have to change to "${ProductId}"-->
      <a href="${}">
        <img src="${}">
      </a>
      <h4 id=productname> "${}" </h4>
      <p id=productdesc> "${}"</p><br>
      <p id=price>"${}"</p>
    </div>
`
document.getElementById(“container”).innerHTML = html;