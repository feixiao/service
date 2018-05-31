// Main.cpp
#include <gtest/gtest.h>


/**
 * 
 *	./foo_test 没有指定filter，运行所有测试；
 *	./foo_test --gtest_filter=* 指定filter为*，运行所有测试；
 *	./foo_test --gtest_filter=FooTest.* 运行测试用例FooTest的所有测试；
 * 	./foo_test --gtest_filter=*Null*:*Constructor* 运行所有全名（即测试用例名 + “ . ” + 测试名，如 GlobalConfigurationTest.noConfigureFileTest）含有"Null"或"Constructor"的测试；
 *  ./foo_test --gtest_filter=FooTest.*-FooTest.Bar 运行测试用例FooTest的所有测试，但不包括FooTest.Bar。
**/

int main(int argc, char** argv) {
	testing::InitGoogleTest(&argc, argv);
	// Runs all tests using Google Test.
	return RUN_ALL_TESTS();
}
