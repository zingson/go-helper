package com.example.demo;

import cn.hutool.core.util.CharsetUtil;
import cn.hutool.crypto.symmetric.SymmetricAlgorithm;
import cn.hutool.crypto.symmetric.SymmetricCrypto;

import java.net.URLEncoder;

public class demo {

    /**
     * 引入工具包
     *         <dependency>
     *             <groupId>cn.hutool</groupId>
     *             <artifactId>hutool-crypto</artifactId>
     *             <version>5.5.4</version>
     *         </dependency>
     * @param args
     */
    public static void main(String[] args) {
        //参数path
        String path = "/二维码页面地址";
        path = path + "/" + "票卡编号";
        //aes加密工具
        String aesKey = "868971231616403394817a2360c4e8b2";
        SymmetricCrypto aes = new SymmetricCrypto(SymmetricAlgorithm.AES, aesKey.getBytes());
        //参数sign，Base64编码
        String sign = aes.encryptBase64("13967119967");
        //url中的参数需要urlEncoder
        String h5Url = "https://itapdev.ucitymetro.com/appentry" + "?" + "path=" + URLEncoder.encode(path, CharsetUtil.CHARSET_UTF_8) +
                "&sign=" + URLEncoder.encode(sign, CharsetUtil.CHARSET_UTF_8);
        System.out.println(h5Url);
    }
}
