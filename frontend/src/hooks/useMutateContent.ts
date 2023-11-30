import { QueryClient, useMutation, useQueryClient } from "@tanstack/react-query"
import useStore from "@/store"
import { ContentType } from "@/pages/api/model/content_model"
import { contentFactory } from "@/pages/api/factory/content_factory"

export const useMutateContent= () => {
    const queryClient = useQueryClient()
    const resetEditedContent = useStore((state) => state.resetContent)

    const createContentMutation = useMutation(
        (content: Omit<ContentType, 'id' | 'createdAt' | 'updatedAt'>) => 
        contentFactory().post(content),
        {
            onSuccess: (res) => {
                const previousContent = queryClient.getQueryData<ContentType[]>([
                    'contents',
                ])
                if (previousContent) {
                    queryClient.setQueryData(
                        ['contents'],
                        [...previousContent, res.data]
                    )
                }
                resetEditedContent()
            },
            onError: (err) => {
                console.log(err)
            }
        }
    )

    const updateContentMutation = useMutation(
        (content: Omit<ContentType, 'createdAt' | 'updatedAt'>) => 
        contentFactory().update(content),
        {
            onSuccess: (res, valiables) => {
                const previousContent = queryClient.getQueryData<ContentType[]>([
                    'contents'
                ])
                if (previousContent) {
                    queryClient.setQueryData<ContentType[]>(
                        ['contents'],
                        previousContent.map((content) =>
                            content.id == valiables.id ? res.data : content
                        )
                    )
                }
            },
            onError: (err) => {
                console.log(err)
            }
        }
    )

    const deleteContentMutation = useMutation(
        (id: number) => contentFactory().delete(id),
        {
            onSuccess: (_, variables) => {
                const previousContent = queryClient.getQueryData<ContentType[]>([
                    'contents'
                ])
                if (previousContent) {
                    queryClient.setQueriesData<ContentType[]>(
                        ['contents'],
                        previousContent.filter((content) => content.id !== variables)
                    )
                }
                resetEditedContent()
            },
            onError: (err) => [
                console.log(err)
            ]
        }
    )

    return {
        createContentMutation,
        updateContentMutation,
        deleteContentMutation
    }
}