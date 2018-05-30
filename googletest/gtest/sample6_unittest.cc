// Copyright 2008 Google Inc.
// All Rights Reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author: wan@google.com (Zhanyong Wan)

// This sample shows how to test common properties of multiple
// implementations of the same interface (aka interface tests).
// 这个示例用来展示如何去测试多个实现了相同接口（类）的共同属性（又叫做接口测试） 

// The interface and its implementations are in this header.
// 接口和他的几个相应实现都在这个头文件中 
#include "prime_tables.h"

#include "gtest/gtest.h"

// First, we define some factory functions for creating instances of
// the implementations.  You may be able to skip this step if all your
// implementations can be constructed the same way.
template <class T>
PrimeTable* CreatePrimeTable();

template <>
PrimeTable* CreatePrimeTable<OnTheFlyPrimeTable>() {
  std::cout<< __FUNCTION__ << std::endl;
  return new OnTheFlyPrimeTable;
}

template <>
PrimeTable* CreatePrimeTable<PreCalculatedPrimeTable>() {
  std::cout<< __FUNCTION__ << std::endl;
  return new PreCalculatedPrimeTable(10000);
}

// Then we define a test fixture class template.
// 接下来我们定义一个test fixture的模版类   
template <class T>
class PrimeTableTest : public testing::Test {
 protected:
  // The ctor calls the factory function to create a prime table
  // implemented by T.
  // 构造函数调用工厂方法通过泛型T创建一个素数表
  PrimeTableTest() : table_(CreatePrimeTable<T>()) {}

  virtual ~PrimeTableTest() { delete table_; }

  // Note that we test an implementation via the base interface
  // instead of the actual implementation class. 
  // 需要注意的是，我们是通过接口而不是他的实现类来测试一个实现的 
  // This is important for keeping the tests close to the real world scenario, where the
  // implementation is invoked via the base interface. 
  // 在那些通过接口来调用实现类的地方，保持测试接近现实世界场景是很重要的 
  //  It avoids got-yas where the implementation class has a method that shadows
  // a method with the same name (but slightly different argument
  // types) in the base interface, for example.
  // 举例来说，在那些实现类中有一个方法名覆盖接口中另一个方法名   

  PrimeTable* const table_;
};

#define GTEST_HAS_TYPED_TEST 1
#define GTEST_HAS_TYPED_TEST_P 1

#if GTEST_HAS_TYPED_TEST

using testing::Types;

// Google Test offers two ways for reusing tests for different types.
// The first is called "typed tests".  You should use it if you
// already know *all* the types you are gonna exercise when you write
// the tests.
// 第一个方法被称作类型测试（typed test）,如果当你写一个测试的时候已经知道了在你的练习中将要用到的所有类型，那么你应该使用这个方法   
// To write a typed test case, first use
//
//   TYPED_TEST_CASE(TestCaseName, TypeList);
//
// to declare it and specify the type parameters.  As with TEST_F,
// TestCaseName must match the test fixture name.
// 正如TEST_F一样，测试案例名字必须和test fixture的名字一致  

// The list of types we want to test.
typedef Types<OnTheFlyPrimeTable, PreCalculatedPrimeTable> Implementations;

// OnTheFlyPrimeTable和PreCalculatedPrimeTable实例都会被创建
TYPED_TEST_CASE(PrimeTableTest, Implementations);

// Then use TYPED_TEST(TestCaseName, TestName) to define a typed test,
// similar to TEST_F.
// 接下来我们像定义TEST_F一样使用TYPED_TEST(TestCaseName, TestName)去定义一个类型测试
TYPED_TEST(PrimeTableTest, ReturnsFalseForNonPrimes) {
  // Inside the test body, you can refer to the type parameter by
  // TypeParam, and refer to the fixture class by TestFixture.  
  // 在测试体中，你可以通过TypeParam指定类型参数，通过TestFixture指定fixture   

  // We  don't need them in this example.
  //在测试体中，你可以通过TypeParam指定类型参数，通过TestFixture指定fixture   
  // Since we are in the template world, C++ requires explicitly
  // writing 'this->' when referring to members of the fixture class.
  // 因为我们都在模版世界之中，在指明fixture类的成员的时候，C++明确要求写'this->' 
  // This is something you have to learn to live with.
  EXPECT_FALSE(this->table_->IsPrime(-5));
  EXPECT_FALSE(this->table_->IsPrime(0));
  EXPECT_FALSE(this->table_->IsPrime(1));
  EXPECT_FALSE(this->table_->IsPrime(4));
  EXPECT_FALSE(this->table_->IsPrime(6));
  EXPECT_FALSE(this->table_->IsPrime(100));
}

TYPED_TEST(PrimeTableTest, ReturnsTrueForPrimes) {
  EXPECT_TRUE(this->table_->IsPrime(2));
  EXPECT_TRUE(this->table_->IsPrime(3));
  EXPECT_TRUE(this->table_->IsPrime(5));
  EXPECT_TRUE(this->table_->IsPrime(7));
  EXPECT_TRUE(this->table_->IsPrime(11));
  EXPECT_TRUE(this->table_->IsPrime(131));
}

TYPED_TEST(PrimeTableTest, CanGetNextPrime) {
  EXPECT_EQ(2, this->table_->GetNextPrime(0));
  EXPECT_EQ(3, this->table_->GetNextPrime(2));
  EXPECT_EQ(5, this->table_->GetNextPrime(3));
  EXPECT_EQ(7, this->table_->GetNextPrime(5));
  EXPECT_EQ(11, this->table_->GetNextPrime(7));
  EXPECT_EQ(131, this->table_->GetNextPrime(128));
}

// That's it!  Google Test will repeat each TYPED_TEST for each type
// in the type list specified in TYPED_TEST_CASE.  Sit back and be
// happy that you don't have to define them multiple times.
// 就是这样，gtest 将会为在TYPED_TEST_CASE指定的类型列表中的每个类型重复一次TYPED_TEST,
// 你不必重复多次地定义他们，你可以感到放松和高兴  

#endif  // GTEST_HAS_TYPED_TEST

#if GTEST_HAS_TYPED_TEST_P

using testing::Types;

// Sometimes, however, you don't yet know all the types that you want
// to test when you write the tests.  For example, if you are the
// author of an interface and expect other people to implement it, you
// might want to write a set of tests to make sure each implementation
// conforms to some basic requirements, but you don't know what
// implementations will be written in the future.
// 然而，有时候当你写测试的时候你也许还不全知道你想测试的所有类型,
// 举例来说，如果你是接口的定义者并且你希望其他的人来实现它,
// 那么你可能想去写一系列的测试去保证每个实现都遵循一些基本要求,
// 但是将来你并不知道那些实现都写了些什么。
//
// How can you write the tests without committing to the type
// parameters?  That's what "type-parameterized tests" can do for you.
// It is a bit more involved than typed tests, but in return you get a
// test pattern that can be reused in many contexts, which is a big
// win.  Here's how you do it:
// First, define a test fixture class template.  Here we just reuse
// the PrimeTableTest fixture defined earlier:
// 如果没有致力于这些类型参数那么你该怎么来写这些测试呢?
// 那就是"type-parameterized tests"（类型参数化测试）能为你做的,他要比类型测试有点复杂。
// 但是反过来你可以得到一个可以在许多情况下重复利用的测试模式，这就是一个收获。
// 这里讲你怎么来做：  
//  首先，定义一个test fixture的模版类(在这里我们重复利用了已经早定义过的PrimeTableTest fixture  )
//    

template <class T>
class PrimeTableTest2 : public PrimeTableTest<T> {
};

// Then, declare the test case.  The argument is the name of the test
// fixture, and also the name of the test case (as usual).  The _P
// suffix is for "parameterized" or "pattern".
// 接下来，声明这个测试案例,参数是这个test fixture的名字，也是测试用例的名字（像平常一样）,
// 这个_P后缀是为了支持“parameterized”（参数化）或者“pattern”（模式）。
TYPED_TEST_CASE_P(PrimeTableTest2);

// Next, use TYPED_TEST_P(TestCaseName, TestName) to define a test,
// similar to what you do with TEST_F.
// 下一步，像你对待TEST_F一样，使用 TYPED_TEST_P(TestCaseName, TestName)来定义你的测试
TYPED_TEST_P(PrimeTableTest2, ReturnsFalseForNonPrimes) {
  EXPECT_FALSE(this->table_->IsPrime(-5));
  EXPECT_FALSE(this->table_->IsPrime(0));
  EXPECT_FALSE(this->table_->IsPrime(1));
  EXPECT_FALSE(this->table_->IsPrime(4));
  EXPECT_FALSE(this->table_->IsPrime(6));
  EXPECT_FALSE(this->table_->IsPrime(100));
}

TYPED_TEST_P(PrimeTableTest2, ReturnsTrueForPrimes) {
  EXPECT_TRUE(this->table_->IsPrime(2));
  EXPECT_TRUE(this->table_->IsPrime(3));
  EXPECT_TRUE(this->table_->IsPrime(5));
  EXPECT_TRUE(this->table_->IsPrime(7));
  EXPECT_TRUE(this->table_->IsPrime(11));
  EXPECT_TRUE(this->table_->IsPrime(131));
}

TYPED_TEST_P(PrimeTableTest2, CanGetNextPrime) {
  EXPECT_EQ(2, this->table_->GetNextPrime(0));
  EXPECT_EQ(3, this->table_->GetNextPrime(2));
  EXPECT_EQ(5, this->table_->GetNextPrime(3));
  EXPECT_EQ(7, this->table_->GetNextPrime(5));
  EXPECT_EQ(11, this->table_->GetNextPrime(7));
  EXPECT_EQ(131, this->table_->GetNextPrime(128));
}

// Type-parameterized tests involve one extra step: you have to
// enumerate the tests you defined:
// 类型参数化测试涉及到一个额外的步骤：你必须要枚举你所定义的测试  
REGISTER_TYPED_TEST_CASE_P(
    //第一个参数是测试案例名
    PrimeTableTest2,  
    // The rest of the arguments are the test names.
    // 其他的是测试名  
    ReturnsFalseForNonPrimes, ReturnsTrueForPrimes, CanGetNextPrime);

// At this point the test pattern is done.  However, you don't have
// any real test yet as you haven't said which types you want to run
// the tests with.
// 这时候测试模式已经完备了,然而，你还没有任何真实的测试，因为你还没有指出使用哪些类型来运行这些测试  

// To turn the abstract test pattern into real tests, you instantiate
// it with a list of types.  Usually the test pattern will be defined
// in a .h file, and anyone can #include and instantiate it.  You can
// even instantiate it more than once in the same program.  To tell
// different instances apart, you give each of them a name, which will
// become part of the test case name and can be used in test filters.
// 为了将这些抽象测试变成真实的测试，你需要用一系列的类型来实例化。 通常，这个测试模式会被定义到一个a.h头文件中，
// 并且任何地方都可以通过#include来示例化他。你甚至可以在同一个程序中不止一次的实例化他，为了区分开实例，你要给每个一个名字。
// 这些名字将会成为测试案例的一部分并且会被应用到测试过滤中 

// The list of types we want to test.  Note that it doesn't have to be
// defined at the time we write the TYPED_TEST_P()s.
// 下面是要测试的类型列表,需要指出的是你在要写TYPED_TEST_P()s时没必要定义这些（参数列表） 
typedef Types<OnTheFlyPrimeTable, PreCalculatedPrimeTable> PrimeTableImplementations;
INSTANTIATE_TYPED_TEST_CASE_P(OnTheFlyAndPreCalculated,    // Instance name  示例名
                              PrimeTableTest2,             // Test case name 测试案例名 
                              PrimeTableImplementations);  // Type list 类型参数

#endif  // GTEST_HAS_TYPED_TEST_P
