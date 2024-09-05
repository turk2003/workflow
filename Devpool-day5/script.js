function showString(stringToShow) {
  document.getElementById("viewData").innerHTML = stringToShow;
}
async function getData() {
  const url =
    "https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_province_with_amphure_tambon.json";

  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    let respData = json;

    // let bkk = respData.filter((elm) => {
    //   return (elm.name_en = 'Bangkok');
    // });
    //console.log(bkk);
    return json;
  } catch (error) {
    console.error(error.message);
  }
}
/* 
  หา list อำเภอ จาก zip_code
*/
function findTumbonFromZipCode(thaiData, zip_code) {
  //thaiData.forEach()
  let arrayOfTumbon = ["1", "2"];
  console.log(arrayOfTumbon.toString());
  return arrayOfTumbon;
}

async function main() {
  let thaiProvinceData = await getData();
  //console.log(thaiProvinceData);
  showString(findTumbonFromZipCode(thaiProvinceData).toString());
}

main();
