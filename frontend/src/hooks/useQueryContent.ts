import { contentFactory } from "@/pages/api/factory/content_factory"
import { ContentType } from "@/pages/api/model/content_model"
import { useQuery } from "@tanstack/react-query"
import { useState } from "react"

export const useQueryContents = () => {
    const [contents, setContents] = useState<ContentType[]>([])
    const getContents = async () => {
        const data = contentFactory().index()
        return data
    }

    return useQuery<ContentType[], Error>({
        queryKey: ['contents'],
        queryFn: getContents,
        staleTime: Infinity,
        onError: (err: any) => {
            if (err.response.data.message) {
                console.log(err.response.data.message)
            } else {
                console.log(err.response.data)
            }
        }
    })
}