<template>
    <p>使用大括号的文本插值：{{ rawHtml }}</p>
    <p>使用 v-html 指令：<span v-html="rawHtml"></span></p>
    <lay-upload @done="getUploadFile" @choose="beginChoose">
        <template #preview>
            <div v-for="(item,index) in picList" :key="`demo1-pic-'${index}`">
                <img :src="item"/>
            </div>
        </template>
    </lay-upload>
</template>

<script>
    import { ref } from 'vue'
    export default {
        name: 'LayuiOne',
        props: {
            message:String
        },
        data(){
            return {
                rawHtml:'<span style="color:red">这里是红色</span>'
            }
        },
        setup() {
            const picList = ref([]);
            const filetoDataURL=(file,fn)=>{
                const reader = new FileReader();
                reader.onloadend = function(e){
                    fn(e.target.result);
                };
                reader.readAsDataURL(file);
            };
            const getUploadFile=(files)=>{
                if(Array.isArray(files)&&files.length>0){
                    files.forEach((file,index,array)=>{
                        filetoDataURL(file,(res)=>{
                            console.log(res);
                            console.log(index);
                            console.log(array);
                            picList.value.push(res);
                            console.log(picList.value);
                        });
                    });
                }
            };
            const beginChoose =(e)=>{
                console.log("beginChoose",e);
            };
            return {
                getUploadFile,
                filetoDataURL,
                beginChoose,
                picList
            }
        }
    }
</script>

