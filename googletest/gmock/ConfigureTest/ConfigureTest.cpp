// ConfigureTest.cpp
#include <gtest/gtest.h>

#include "Configure.h"

TEST(ConfigureTest, addItem)
{
	// do some initialization
	Configure* pc = new Configure();
	
	// validate the pointer is not null
	ASSERT_TRUE(pc != NULL);

	// call the method we want to test
	pc->addItem("A");
	pc->addItem("B");
	pc->addItem("A");

	// validate the result after operation
	EXPECT_EQ(pc->getSize(), 2);
	EXPECT_STREQ(pc->getItem(0).c_str(), "A");
	EXPECT_STREQ(pc->getItem(1).c_str(), "B");
	EXPECT_STREQ(pc->getItem(10).c_str(), "");

	delete pc;
}
