var a = 1;

var items = [];

function getItem()
{
    var item = document.getElementById('input');
    item = item.value;

    createItem(item);
    render();
}

function createItem(item)
{
    var id = '' + new Date().getTime()

    items.push({
        id: id,
        name: item 
    });
    console.log(item.name)
}

function render()
{
    document.getElementById('item-list').innerHTML = '';
    items.forEach(function(item)
    {
        var element = document.createElement('div')
        element.innerText = '' + item.name;
        var itemlist = document.getElementById('item-list');
        itemlist.appendChild(element);
    });
}

function putItem(item)
{

}