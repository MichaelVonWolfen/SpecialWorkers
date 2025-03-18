package main

import (
	"SpecialWorkers/models"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
)

func worker(data models.WorkerInformation) []models.MangaInformation {
	return nil
}
func preProcess(raw string) string {
	//Cleaning string from the serialized placeable locations
	re := regexp.MustCompile(`"\$L\d+"`)
	clean := re.ReplaceAllString(raw, "null")

	//extract all arrays
	beginning := strings.Index(clean, "[")
	if beginning == -1 {
		log.Fatal("No arraysQ")
	}
	return clean[beginning:]
}

func walkTree(node interface{}, extracted *[]map[string]interface{}) {
	//var dataList map[string]interface{}
	var data = make(map[string]interface{})
	switch n := node.(type) {
	case []interface{}:
		if len(n) > 1 {
			if tag, ok := n[1].(string); ok {
				//fmt.Printf("Found tag: %s\n", tag)
				switch tag {
				case "a":
					if len(n) > 3 {
						if props, ok := n[3].(map[string]interface{}); ok {
							if val, ok := props["href"]; ok {
								href := val
								//fmt.Printf("href: %s\n", href)
								data["url"] = href
							}
						}
					}
				case "span":
					if len(n) > 3 {
						if props, ok := n[3].(map[string]interface{}); ok {
							if val, ok := props["children"]; ok {
								if children, ok := val.(string); ok {
									//fmt.Printf("children: %s\n", val)
									//fmt.Println(children)
									if slices.Contains([]string{"Ongoing", "Completed", "Hiatus", "Cancelled", "Unknown"}, children) {
										data["status"] = children
									} else {
										//data["status"] = "Unknown"
										data["title"] = children
									}
								}
							}
						}
					}
				}
			}
		}
		for _, item := range n {
			walkTree(item, extracted)
		}
	case map[string]interface{}:
		for _, v := range n {
			walkTree(v, extracted)
		}
	}
	if _, ok := data["url"]; ok {
		*extracted = append(*extracted, data)
	}
	if _, ok := data["status"]; ok {
		*extracted = append(*extracted, data)
	}
	if _, ok := data["title"]; ok {
		*extracted = append(*extracted, data)
	}
}
func processData(rawStr string) {
	str := preProcess(rawStr)
	var root interface{}
	if err := json.Unmarshal([]byte(str), &root); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		return
	}

	var extracted []map[string]interface{}
	walkTree(root, &extracted)

	for _, item := range extracted {
		fmt.Println(item)
	}
	fmt.Println(extracted)
	//println(str)
}
func main() {
	var testAsura = models.WorkerInformation{
		WebsiteSearchUrl: "https://api.reaperscans.com/query?adult=true&query_string=",
		MangaName:        "Solo",
		MangaId:          0,
		//mangaUrl:             "https://api.reaperscans.com/chapters/100?page=1&perPage=1000&query=&order=desc",
		WebsiteSearchPattern: "",
		MangaSearchPattern:   "",
		ChapterList:          nil,
	}
	worker(testAsura)
	i := 0
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`3:[["$","meta","0",{"name":"viewport","content":"width=device-width, initial-scale=1"}],["$","meta","1",{"charSet":"utf-8"}],["$","title","2",{"children":"Series - Asura Scans"}],["$","meta","3",{"name":"description","content":"Explore the latest manga on Asura Scans."}],["$","meta","4",{"property":"og:title","content":"Manga - Asura Scans"}],["$","meta","5",{"property":"og:description","content":"Explore the latest manga on Asura Scans."}],["$","meta","6",{"property":"og:url","content":"https://asuracomic.net/series/"}],["$","meta","7",{"property":"og:site_name","content":"Asura Scans"}],["$","meta","8",{"property":"og:locale","content":"en_US"}],["$","meta","9",{"property":"og:type","content":"website"}],["$","meta","10",{"name":"twitter:card","content":"summary_large_image"}],["$","meta","11",{"name":"twitter:title","content":"Manga - Asura Scans"}],["$","meta","12",{"name":"twitter:description","content":"Explore the latest manga on Asura Scans."}],["$","meta","13",{"name":"next-size-adjust"}]]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`4:I[74733,["8791","static/chunks/08ffd5a1-79677cef068ed986.js","1304","static/chunks/1304-851c1a7e9060ddca.js","3043","static/chunks/3043-8f89080a8f6b6d23.js","7200","static/chunks/7200-a4a9bb38f630a460.js","6546","static/chunks/6546-e7ce7ba14de4a620.js","4583","static/chunks/4583-279896ab4aefae78.js","4911","static/chunks/4911-1cf42893664a505f.js","8731","static/chunks/8731-19787f6d128b358b.js","5936","static/chunks/5936-c949c38ce30a7cab.js","6282","static/chunks/6282-479b8146bba17175.js","4733","static/chunks/4733-ae1d629798f70661.js","2989","static/chunks/app/(comic)/series/(page)/page-a67ce800095af24b.js"],"default"]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`5:I[67200,["8791","static/chunks/08ffd5a1-79677cef068ed986.js","1304","static/chunks/1304-851c1a7e9060ddca.js","3043","static/chunks/3043-8f89080a8f6b6d23.js","7200","static/chunks/7200-a4a9bb38f630a460.js","6546","static/chunks/6546-e7ce7ba14de4a620.js","4583","static/chunks/4583-279896ab4aefae78.js","4911","static/chunks/4911-1cf42893664a505f.js","8731","static/chunks/8731-19787f6d128b358b.js","5936","static/chunks/5936-c949c38ce30a7cab.js","6282","static/chunks/6282-479b8146bba17175.js","4733","static/chunks/4733-ae1d629798f70661.js","2989","static/chunks/app/(comic)/series/(page)/page-a67ce800095af24b.js"],"Image"]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`6:I[38731,["8791","static/chunks/08ffd5a1-79677cef068ed986.js","1304","static/chunks/1304-851c1a7e9060ddca.js","3043","static/chunks/3043-8f89080a8f6b6d23.js","7200","static/chunks/7200-a4a9bb38f630a460.js","6546","static/chunks/6546-e7ce7ba14de4a620.js","4583","static/chunks/4583-279896ab4aefae78.js","4911","static/chunks/4911-1cf42893664a505f.js","8731","static/chunks/8731-19787f6d128b358b.js","5936","static/chunks/5936-c949c38ce30a7cab.js","6282","static/chunks/6282-479b8146bba17175.js","4733","static/chunks/4733-ae1d629798f70661.js","2989","static/chunks/app/(comic)/series/(page)/page-a67ce800095af24b.js"],"Rating"]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`7:I[38731,["8791","static/chunks/08ffd5a1-79677cef068ed986.js","1304","static/chunks/1304-851c1a7e9060ddca.js","3043","static/chunks/3043-8f89080a8f6b6d23.js","7200","static/chunks/7200-a4a9bb38f630a460.js","6546","static/chunks/6546-e7ce7ba14de4a620.js","4583","static/chunks/4583-279896ab4aefae78.js","4911","static/chunks/4911-1cf42893664a505f.js","8731","static/chunks/8731-19787f6d128b358b.js","5936","static/chunks/5936-c949c38ce30a7cab.js","6282","static/chunks/6282-479b8146bba17175.js","4733","static/chunks/4733-ae1d629798f70661.js","2989","static/chunks/app/(comic)/series/(page)/page-a67ce800095af24b.js"],"RoundedStar"]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`8:I[58346,["8791","static/chunks/08ffd5a1-79677cef068ed986.js","1304","static/chunks/1304-851c1a7e9060ddca.js","3043","static/chunks/3043-8f89080a8f6b6d23.js","7200","static/chunks/7200-a4a9bb38f630a460.js","6546","static/chunks/6546-e7ce7ba14de4a620.js","4583","static/chunks/4583-279896ab4aefae78.js","4911","static/chunks/4911-1cf42893664a505f.js","8731","static/chunks/8731-19787f6d128b358b.js","5936","static/chunks/5936-c949c38ce30a7cab.js","6282","static/chunks/6282-479b8146bba17175.js","4733","static/chunks/4733-ae1d629798f70661.js","2989","static/chunks/app/(comic)/series/(page)/page-a67ce800095af24b.js"],"default"]`)
	i += 1
	fmt.Printf("\t\tWorking on entry nb %d\n", i)
	processData(`2:["$","div",null,{"className":"w-full min-[768px]:w-[100%] bg-[#222222] min-[880px]:w-[98%] min-[912px]:w-[98%] lg:w-[100%] mb-2","children":[["$","div",null,{"className":"relative flex justify-between align-baseline font-500 bg-[#222222] border-b-[1px] border-[#312f40] px-[15px] py-[8px] p-2  mt-5 sm:mt-0","children":["$","h3",null,{"className":"text-[15px] text-white font-semibold leading-5 m-0","children":"Series list"}]}],["$","$L4",null,{"params":{"page":"1","name":"solo"},"genres":"$undefined","paramsLength":2,"type":"series"}],["$","div",null,{"className":"grid grid-cols-2 sm:grid-cols-2 md:grid-cols-5 gap-3 p-4","children":[false,[["$","a","0",{"href":"series/solo-max-level-newbie-4c3ee20c","children":["$","div",null,{"className":"w-full block sm:block hover:cursor-pointer group hover:text-themecolor","children":["$","div",null,{"children":[["$","div",null,{"className":"flex h-[250px] md:h-[200px] overflow-hidden relative hover:opacity-60","children":[["$","span",null,{"className":"status bg-blue-700","children":"Ongoing"}],["$","$L5",null,{"src":"https://gg.asuracomic.net/storage/media/272496/conversions/01JMHFP0DBPD906JMCZNAKG1RH-thumb-small.webp","alt":"","width":0,"height":0,"sizes":"100vh","style":{"width":"100%","height":"100%","objectFit":"cover","objectPosition":"top"},"className":"rounded-md"}],["$","div",null,{"className":"absolute bottom-[0px] flex justify-center left-[5px] mb-[5px] rounded-[3px] text-white bg-[#a12e24] ","children":["$","span",null,{"className":"text-[10px] font-bold py-[2px] px-[7px]","children":"MANHWA"}]}]]}],["$","div",null,{"className":"block w-[100%] h-auto  items-center","style":{"fontSize":"13.3px","margin":"8px 0","marginBottom":"5px","fontWeight":500,"lineHeight":"19px","textAlign":"left","overflow":"hidden"},"children":[["$","span",null,{"className":"block text-[13.3px] font-bold","children":"Solo Max-Level Newbie"}],["$","span",null,{"className":"text-[13px] text-[#999]","children":["Chapter ",197]}],["$","span",null,{"className":"flex text-[12px] text-[#999]","children":[["$","$L6",null,{"style":{"maxWidth":70},"value":4.95,"readOnly":true,"itemStyles":{"itemShapes":"$7","activeFillColor":"#ffc700","inactiveFillColor":"#686868"},"items":5}],["$","label",null,{"className":"ml-1","children":9.9}]]}]]}]]}]}]}],["$","a","1",{"href":"series/solo-leveling-ragnarok-a9340788","children":["$","div",null,{"className":"w-full block sm:block hover:cursor-pointer group hover:text-themecolor","children":["$","div",null,{"children":[["$","div",null,{"className":"flex h-[250px] md:h-[200px] overflow-hidden relative hover:opacity-60","children":[["$","span",null,{"className":"status bg-blue-700","children":"Ongoing"}],["$","$L5",null,{"src":"https://gg.asuracomic.net/storage/media/271/conversions/01J3QRQXHSBQVGA2KHKP2D3S3S-thumb-small.webp","alt":"","width":0,"height":0,"sizes":"100vh","style":{"width":"100%","height":"100%","objectFit":"cover","objectPosition":"top"},"className":"rounded-md"}],["$","div",null,{"className":"absolute bottom-[0px] flex justify-center left-[5px] mb-[5px] rounded-[3px] text-white bg-[#a12e24] ","children":["$","span",null,{"className":"text-[10px] font-bold py-[2px] px-[7px]","children":"MANHWA"}]}]]}],["$","div",null,{"className":"block w-[100%] h-auto  items-center","style":{"fontSize":"13.3px","margin":"8px 0","marginBottom":"5px","fontWeight":500,"lineHeight":"19px","textAlign":"left","overflow":"hidden"},"children":[["$","span",null,{"className":"block text-[13.3px] font-bold","children":"Solo Leveling: Ragnarok"}],["$","span",null,{"className":"text-[13px] text-[#999]","children":["Chapter ",45]}],["$","span",null,{"className":"flex text-[12px] text-[#999]","children":[["$","$L6",null,{"style":{"maxWidth":70},"value":4.85,"readOnly":true,"itemStyles":{"itemShapes":"$7","activeFillColor":"#ffc700","inactiveFillColor":"#686868"},"items":5}],["$","label",null,{"className":"ml-1","children":9.7}]]}]]}]]}]}]}],["$","a","2",{"href":"series/solo-spell-caster-af1398af","children":["$","div",null,{"className":"w-full block sm:block hover:cursor-pointer group hover:text-themecolor","children":["$","div",null,{"children":[["$","div",null,{"className":"flex h-[250px] md:h-[200px] overflow-hidden relative hover:opacity-60","children":[["$","span",null,{"className":"status bg-[#de3b3b]","children":"Dropped"}],["$","$L5",null,{"src":"https://gg.asuracomic.net/storage/media/180/conversions/da69f53a-thumb-small.webp","alt":"","width":0,"height":0,"sizes":"100vh","style":{"width":"100%","height":"100%","objectFit":"cover","objectPosition":"top"},"className":"rounded-md"}],["$","div",null,{"className":"absolute bottom-[0px] flex justify-center left-[5px] mb-[5px] rounded-[3px] text-white bg-[#a12e24] ","children":["$","span",null,{"className":"text-[10px] font-bold py-[2px] px-[7px]","children":"MANHWA"}]}]]}],["$","div",null,{"className":"block w-[100%] h-auto  items-center","style":{"fontSize":"13.3px","margin":"8px 0","marginBottom":"5px","fontWeight":500,"lineHeight":"19px","textAlign":"left","overflow":"hidden"},"children":[["$","span",null,{"className":"block text-[13.3px] font-bold","children":"Solo Spell Caster"}],["$","span",null,{"className":"text-[13px] text-[#999]","children":["Chapter ",87]}],["$","span",null,{"className":"flex text-[12px] text-[#999]","children":[["$","$L6",null,{"style":{"maxWidth":70},"value":4.15,"readOnly":true,"itemStyles":{"itemShapes":"$7","activeFillColor":"#ffc700","inactiveFillColor":"#686868"},"items":5}],["$","label",null,{"className":"ml-1","children":8.3}]]}]]}]]}]}]}],["$","a","3",{"href":"series/solo-bug-player-603704d7","children":["$","div",null,{"className":"w-full block sm:block hover:cursor-pointer group hover:text-themecolor","children":["$","div",null,{"children":[["$","div",null,{"className":"flex h-[250px] md:h-[200px] overflow-hidden relative hover:opacity-60","children":[["$","span",null,{"className":"status bg-[#de3b3b]","children":"Dropped"}],["$","$L5",null,{"src":"https://gg.asuracomic.net/storage/media/245/01J3BAR5EFJJSB84FC5GDZYSW7.webp","alt":"","width":0,"height":0,"sizes":"100vh","style":{"width":"100%","height":"100%","objectFit":"cover","objectPosition":"top"},"className":"rounded-md"}],["$","div",null,{"className":"absolute bottom-[0px] flex justify-center left-[5px] mb-[5px] rounded-[3px] text-white bg-[#a12e24] ","children":["$","span",null,{"className":"text-[10px] font-bold py-[2px] px-[7px]","children":"MANHWA"}]}]]}],["$","div",null,{"className":"block w-[100%] h-auto  items-center","style":{"fontSize":"13.3px","margin":"8px 0","marginBottom":"5px","fontWeight":500,"lineHeight":"19px","textAlign":"left","overflow":"hidden"},"children":[["$","span",null,{"className":"block text-[13.3px] font-bold","children":"Solo Bug Player"}],["$","span",null,{"className":"text-[13px] text-[#999]","children":["Chapter ",88]}],["$","span",null,{"className":"flex text-[12px] text-[#999]","children":[["$","$L6",null,{"style":{"maxWidth":70},"value":4.35,"readOnly":true,"itemStyles":{"itemShapes":"$7","activeFillColor":"#ffc700","inactiveFillColor":"#686868"},"items":5}],["$","label",null,{"className":"ml-1","children":8.7}]]}]]}]]}]}]}],["$","a","4",{"href":"series/solo-leveling-a89d6a6f","children":["$","div",null,{"className":"w-full block sm:block hover:cursor-pointer group hover:text-themecolor","children":["$","div",null,{"children":[["$","div",null,{"className":"flex h-[250px] md:h-[200px] overflow-hidden relative hover:opacity-60","children":[["$","span",null,{"className":"status bg-[#de3b3b]","children":"Completed"}],["$","$L5",null,{"src":"https://gg.asuracomic.net/storage/media/256/conversions/01J3BAXFBTABT3VNAV3RPNZK7S-thumb-small.webp","alt":"","width":0,"height":0,"sizes":"100vh","style":{"width":"100%","height":"100%","objectFit":"cover","objectPosition":"top"},"className":"rounded-md"}],["$","div",null,{"className":"absolute bottom-[0px] flex justify-center left-[5px] mb-[5px] rounded-[3px] text-white bg-[#a12e24] ","children":["$","span",null,{"className":"text-[10px] font-bold py-[2px] px-[7px]","children":"MANHWA"}]}]]}],["$","div",null,{"className":"block w-[100%] h-auto  items-center","style":{"fontSize":"13.3px","margin":"8px 0","marginBottom":"5px","fontWeight":500,"lineHeight":"19px","textAlign":"left","overflow":"hidden"},"children":[["$","span",null,{"className":"block text-[13.3px] font-bold","children":"Solo Leveling"}],["$","span",null,{"className":"text-[13px] text-[#999]","children":["Chapter ",200]}],["$","span",null,{"className":"flex text-[12px] text-[#999]","children":[["$","$L6",null,{"style":{"maxWidth":70},"value":5,"readOnly":true,"itemStyles":{"itemShapes":"$7","activeFillColor":"#ffc700","inactiveFillColor":"#686868"},"items":5}],["$","label",null,{"className":"ml-1","children":10}]]}]]}]]}]}]}]]]}],["$","div",null,{"className":"flex items-center justify-center py-[15px] bg-[#222222] ","children":["$","$L8",null,{"searchParams":"$9","prevPageUrl":null,"currentPage":1,"nextPageUrl":null}]}]]}]`)

	//TODO: SEND RETRIEVED DATA TO THE DB???
}
