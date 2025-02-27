
use serde::{Deserialize, Serialize};
use serde_json;



#[derive(Serialize, Deserialize, Debug)]
struct Person {
    items: Vec<u32>,
    name: String,
    phone: Phone
}

#[derive(Serialize, Deserialize, Debug)]
struct Phone {
    location: String,
    number: String,
}

fn main() {

    // == Declare a map or HashMap
    // let mut map : HashMap<&str, serde_json::Value>= HashMap::new();
    let mut map = serde_json::Map::new();
    // better to use json! to convert something into a JsonValue
    map.insert("items".to_string(), serde_json::to_value(vec![1,2,3,4,5]).unwrap());
    map.insert("name".to_string(), serde_json::to_value("Mike").unwrap());
    map.insert("phone".to_string(), serde_json::json!(Phone {
        location:("New York").to_string(),
        number:("+123456").to_string(),
    })) ;

    // == 1. Both ways can convert a map into a JsonValue
    // let json = serde_json::to_value(&map).unwrap();
    let  json = serde_json::Value::Object(map);

    // == Convert a map into a json string
    // let json_string = serde_json::to_string(&map).unwrap();

    // == 2. Convert a JsonValue into map
    if let Some(map2) = json.as_object() {
    println!("{:?}", map2);
    }


    // == 3. Convert a struct into JsonValue
    let person = Person{
        items: vec![1,2,3,],
        name: "John".to_string(),
        phone: Phone {
            location: "Chicago".to_string(),
            number: "+3322323".to_string(),
        }
    };

    let json_person = serde_json::to_value(&person).unwrap();

    // == 4. Convert a JsonValue to struct
    let person2: Person = serde_json::from_value(json_person).unwrap();


}


