import { ContentType } from "../model/content_model";
import { ContentRepository, contentRepository } from "../repositories/content_repositories";

// mockを入れることも可能
export const contentFactory = (req?: ContentRepository) => {
    const repository = req ?? contentRepository
    return {
        index : async (): Promise<ContentType[]> => {
            const response = await repository.getContents()
            return response
        },
        show: async (id: Pick<ContentType, 'id'>): Promise<ContentType> => {
            const response = await repository.getContent(id)
            return response
        },
        post: async (
            params: Omit<ContentType, 'id' | 'createdAt' | 'updatedAt'>
        ) => {
            const response = await repository.createContent(params)
            return response
        },
        update: async (
            content: Omit<ContentType, 'createdAt' | 'updatedAt'>
        ) => {
            const response = await repository.updateContent(content)
            return response
        },
        delete: async (id: number) => {
            const response = await repository.deleteContent(id)
            return response
        }
    }
}