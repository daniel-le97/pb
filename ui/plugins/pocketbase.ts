import type { TypedPocketBase } from 'typed-pocketbase';
import type { Schema } from '../Database';

export default defineNuxtPlugin(async nuxt => {
    const PocketBase = (await import('pocketbase')).default
    const pb = new PocketBase( "http://localhost:8090" ) as TypedPocketBase<Schema>

    return {
        provide:{
            pb
        }
    }

})