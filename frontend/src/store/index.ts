import {create} from 'zustand'

type EditedContent = {
    id : number,
    title: string
}

type State = {
    editedContent: EditedContent,
    updateContent: (payload: EditedContent) => void,
    resetContent: () => void
}

const useStore = create<State>((set) => ({
    editedContent: { id: 0, title: '' },
    updateContent: (payload) => {
        set({
            editedContent: payload
        })
    },
    resetContent: () => set({editedContent: { id: 0, title: ''}})
}))

export default useStore