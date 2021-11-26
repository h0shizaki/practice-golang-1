function app(){

    

    fetch("http://localhost:3000/data")
    .then(res => res.json())
    .then( data=>{
        let text= "";

        data.forEach( (user) =>{
            text +=  "<li>" +user.id+" : "+ user.name + " </li>"
        })
        console.log(text)
    
        document.querySelector("#list").innerHTML = text
    })


 
}
app()