import { writable } from "svelte/store"

const cache = new Map()
export const mcompany = writable([])
export const fetchCompany = async(url) => {
    if(cache.has(url)){
        console.log("cache")
        mcompany.set(cache.get(url))
    }else{
        console.log("dari url")
        const res = await fetch(url)
        const data = await res.json()
        const loadedCompany = data.record.map((data, index) => {
            return {
                no : index + 1,
                idcompany: data.idcompany,
                idcurr: data.idcurr,
                nmcompany: data.nmcompany,
                nmowner: data.nmowner,
                statuscompany: data.statuscompany,
            }
        })
        cache.set(url, loadedCompany)
        mcompany.set(loadedCompany)
    }
}