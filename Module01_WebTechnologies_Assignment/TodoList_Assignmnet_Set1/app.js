function addTask(){
    const task = document.getElementById("input");
    const text = task.value.trim();
    const list = document.getElementById("list");

    const li = document.createElement("li");
    li.textContent = text;

    const btn = document.createElement("button");
    btn.classList.add("btn", "btn-outline-success", "btn-sm");
    btn.id = "btn-done";
    btn.type = "button";
    btn.textContent = "mark as completed";
    btn.onclick = () => markDone(btn);
    li.appendChild(btn);

    const btn_delete = document.createElement("button");
    btn_delete.classList.add("btn-close", "btn", "btn-danger");
    btn_delete.id = "btn-delete";
    btn_delete.type = "button";
    btn_delete.setAttribute("aria-label", "Close");
    btn_delete.onclick = () => deleteTask(btn_delete);
    li.appendChild(btn_delete);
    
    const para = document.getElementById("para");
    if(para != null){
        para.remove();
    }

    list.appendChild(li);
    task.value = "";
}

function markDone(btn){
    console.log(btn);
    const parentListItem = btn.closest("li");
    parentListItem.style.textDecoration = "line-through";
    btn.disabled = true;
    parentListItem.style.color = "grey";
}

function deleteTask(btn_delete){
    const parentListItem = btn_delete.closest("li");
    parentListItem.remove();

    const list = document.getElementById("list");
    if(list.childElementCount === 0){
        const para = document.createElement("p");
        const div = document.getElementById("div")
        para.textContent = "No Tasks added yet...";
        para.id = "para";
        para.style.color = "#a3a8a9bd";
        para.style.fontSize = "30px";
        para.style.margin = "0 auto";
        para.style.fontStyle = "italic";
        div.appendChild(para);
    }
}