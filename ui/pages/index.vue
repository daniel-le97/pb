<template>
    <div v-if=" state ">
        <div v-for=" item  in  state ">
            {{ item }}
        </div>
    </div>
</template>
<script setup>
const state = ref();
state.value = await usePB().collection( 'git' ).getFullList();


const okay = await usePB().collection( 'git' ).subscribe( '*', ( data ) =>
{
    console.log( data );
    if ( data.action === 'create' )
    {
        state.value.push( data.record );
    }
} );
</script>