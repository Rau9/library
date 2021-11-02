$(document).ready(function(){
    // Logic when we are in consult
     if(window.location.href.indexOf('consult') > -1){
        let tabla = document.getElementById("table");
        let info;
         fetch('http://localhost:3000/books',
         {
            method: 'GET',
            mode: 'no-cors'
         }
         ).then( resp => resp.json() )
         .then( datos => {
               info= datos.books.map(e => {
                   return `
                <thead>
                    <tr>
                        <th>ID:</th><td>${e.ID}</td>
                        <th>ISBN:</th><td>${e.ISBN}</td>
                        <th>Título:</th><td>${e.Title}</td>
                        <th>Autor:</th><td>${e.Author.Name}</td>
                        <th>Descripción:</th><td>${e.Description}</td>
                        <th>Categoría:</th><td>${e.Category}</td>
                    </tr>
                 </thead>`
                }).join('');
                tabla.innerHTML = info;
                console.log(info)
            })
            .catch( err => console.log(err) );
    } 
})

// This function would be used to send back the data of the registered user to persist it in the DB
function submitUser(){
    let userName = document.getElementById("userRegister").value;
    let password = document.getElementById("passwordRegister").value;
    let passwordConfig = document.getElementById("passwordRegisterConfirm").value;
    if(password == passwordConfig){
        fetch('url', 
        {
            method: 'POST',
            body:
            JSON.stringify({
                userName : userName,
                password : password
            }),
            headers: {
                "Content-type":"application/json"
            }
        })
        .then(resp=>resp.json())
        .then(data=>console.log(data))
    }else{
        alert("Las contraseñas no coinciden");
    }
      
}

// Logic when we press in init session
function login(){
    if(document.getElementById("loginUser").value == 'admin' && document.getElementById("passwordLogin").value == 'admin'){
        window.location.href = "consult.html";
    }else{ 
        alert("Usuario o contraseña incorrectos"); 
    }
}

// Logic when we press in create
function create(){  
    let isbnCreate = document.getElementById("isbnCreate").value;
    let titleCreate = document.getElementById("titleCreate").value;
    let authorCreate = document.getElementById("authorCreate").value;
    let descriptionCreate = document.getElementById("descriptionCreate").value;
    let categoryCreate = document.getElementById("categoryCreate").value;
    let birthDate = document.getElementById("birthDate").value;
    let birthDead = document.getElementById("birthDead").value;
    let nick = document.getElementById("nick").value;
    if(isbnCreate == "" || titleCreate == "" || authorCreate == "" || descriptionCreate == "" || categoryCreate == ""){
        alert("Introduzca todos los datos obligatorios");
    }else{
        fetch('http://localhost:3000/books', 
        {
            method: 'POST',
            body:
            JSON.stringify({
                isbn : isbnCreate,
                title : titleCreate,
                description : descriptionCreate,
                category : categoryCreate,
                author :{
                            name: authorCreate,
                            date_of_birth: birthDate,
                            date_of_death: birthDead,
                            nick:  nick

                        }                      
            }),
            headers: {
                "Content-type":"application/json"
            }
            })
            .then(function(){
                document.getElementById("isbnCreate").value = "";
                document.getElementById("titleCreate").value = "";
                document.getElementById("authorCreate").value = "";; 
                document.getElementById("descriptionCreate").value = "";
                document.getElementById("categoryCreate").value = "";
                document.getElementById("birthDate").value = "";
                document.getElementById("birthDead").value = "";
                document.getElementById("nick").value = "";
                alert("Libro creado correctamente");
            }) 
        }       

}

// Logic when we press in modify
function modify(){
    let idModify = document.getElementById("idModify").value;
    let isbnModify = document.getElementById("isbnModify").value;
    let titleModify = document.getElementById("titleModify").value;
    let descriptionModify = document.getElementById("descriptionModify").value;
    let category = document.getElementById("category").value;

    if(idModify == "" || isbnModify == "" || titleModify == "" || descriptionModify == "" || category == ""){
        alert("Introduzca todos los datos obligatorios");
    }else{
    
        fetch('http://localhost:3000/books/' + idModify, 
        {
            method: 'PUT',
            body:
            JSON.stringify({
                isbn : isbnModify,
                title : titleModify,
                description : descriptionModify,
                category : category          
            }),
            headers: {
                "Content-type":"application/json"
            }
        })
        .then(function(){
            document.getElementById("idModify").value = "";
            document.getElementById("isbnModify").value = "";
            document.getElementById("titleModify").value = "";
            document.getElementById("descriptionModify").value = "";
            document.getElementById("category").value = "";
            alert("Libro modificado correctamente");
        })
    }
        
}

// Logic when press in remove
 function remove(){
    let idRemove = document.getElementById("idnRemove").value;

    if(idRemove == ""){
        alert("Introduzca un ID");
    }else{
        fetch('http://localhost:3000/books/'+ idRemove,
        {
            method:'DELETE'
        })
        .then(function(){
            document.getElementById("idnRemove").value = "";
            alert("Libro eliminado correctamente");
        })
    }     
 }
