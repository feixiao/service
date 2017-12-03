package com.frank;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Log {

    Logger logger = LoggerFactory.getLogger(getClass());

    public static void main(String[] args) {
        Log log = new Log();
        log.test();
    }


    public void test() {
        logger.debug("debug");
        logger.info("info");
    }
}
