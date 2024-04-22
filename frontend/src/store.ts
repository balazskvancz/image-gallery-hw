import { writable } from 'svelte/store'

export const onRefreshData = writable<boolean>(false)