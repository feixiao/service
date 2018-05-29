// Copyright 2005, Google Inc.
// All rights reserved.
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

// A sample program demonstrating using Google C++ testing framework.
//
// Author: wan@google.com (Zhanyong Wan)


// In this example, we use a more advanced feature of Google Test called
// test fixture.
//
// A test fixture is a place to hold objects and functions shared by
// all tests in a test case.  Using a test fixture avoids duplicating
// the test code necessary to initialize and cleanup those common
// objects for each test.  It is also useful for defining sub-routines
// that your tests need to invoke a lot.
//
// <TechnicalDetails>
//
// The tests share the test fixture in the sense of code sharing, not
// data sharing.  Each test is given its own fresh copy of the
// fixture.  You cannot expect the data modified by one test to be
// passed on to another test, which is a bad idea.
//
// The reason for this design is that tests should be independent and
// repeatable.  In particular, a test should not fail as the result of
// another test's failure.  If one test depends on info produced by
// another test, then the two tests should really be one big test.
//
// The macros for indicating the success/failure of a test
// (EXPECT_TRUE, FAIL, etc) need to know what the current test is
// (when Google Test prints the test result, it tells you which test
// each failure belongs to).  Technically, these macros invoke a
// member function of the Test class.  Therefore, you cannot use them
// in a global function.  That's why you should put test sub-routines
// in a test fixture.
//
// </TechnicalDetails>

#include "sample3-inl.h"
#include "gtest/gtest.h"


// 当你发现自己编写了两个或多个测试来操作同样的数据，你可以采用一个测试固件。
// 它让你可以在多个不同的测试中重用同样的对象配置。

// 创建测试固件的步骤
// 1: 创建一个类继承自testing::Test。将其中的成员声明为protected:或是public:，因为我们想要从子类中存取固件成员。
// 2: 在该类中声明你计划使用的任何对象。
// 3: 如果需要，编写一个默认构造函数或者SetUp()函数来为每个测试准备对象。常见错误包括将SetUp()拼写为Setup()（小写了u）——不要让它发生在你身上。
// 4: 如果需要，编写一个析构函数或者TearDown()函数来释放你在SetUp()函数中申请的资源。要知道什么时候应该使用构造函数/析构函数，什么时候又应该使用SetUp()/TearDown()函数，阅读我们的FAQ。
// 5: 如果需要，定义你的测试所需要共享的子程序。

// 注：当我们要使用固件时，使用TEST_F()替换掉TEST()
// 对于TEST_F()中定义的每个测试，Google Test将会：
//    1: 在运行时创建一个全新的测试固件
//    2: 马上通过SetUp()初始化它，
//    3: 运行测试
//    4: 调用TearDown()来进行清理工作
//    5: 删除测试固件。
//    注意，同一测试案例中，不同的测试拥有不同的测试固件。Google Test在创建下一个测试固件前总是会对现有固件进行删除。
//    Google Test不会对多个测试重用一个测试固件。测试对测试固件的改动并不会影响到其他测试。

// To use a test fixture, derive a class from testing::Test.
class QueueTest : public testing::Test {
 protected:  // You should make the members protected s.t. they can be
             // accessed from sub-classes.

  // virtual void SetUp() will be called before each test is run.  You
  // should define it if you need to initialize the varaibles.
  // Otherwise, this can be skipped.
  virtual void SetUp() {
    q1_.Enqueue(1);
    q2_.Enqueue(2);
    q2_.Enqueue(3);
  }

  // virtual void TearDown() will be called after each test is run.
  // You should define it if there is cleanup work to do.  Otherwise,
  // you don't have to provide it.
  //
  // virtual void TearDown() {
  // }

  // A helper function that some test uses.
  static int Double(int n) {
    return 2*n;
  }

  // A helper function for testing Queue::Map().
  void MapTester(const Queue<int> * q) {
    // Creates a new queue, where each element is twice as big as the
    // corresponding one in q.
    const Queue<int> * const new_q = q->Map(Double);

    // Verifies that the new queue has the same size as q.
    ASSERT_EQ(q->Size(), new_q->Size());

    // Verifies the relationship between the elements of the two queues.
    for ( const QueueNode<int> * n1 = q->Head(), * n2 = new_q->Head();
          n1 != NULL; n1 = n1->next(), n2 = n2->next() ) {
      EXPECT_EQ(2 * n1->element(), n2->element());
    }

    delete new_q;
  }

  // Declares the variables your tests want to use.
  Queue<int> q0_;
  Queue<int> q1_;
  Queue<int> q2_;
};

// When you have a test fixture, you define a test using TEST_F
// instead of TEST.

// Tests the default c'tor.
TEST_F(QueueTest, DefaultConstructor) {
  // You can access data in the test fixture here.
  // Setup中 q0_ 没有添加数据，所以测试没有问题
  EXPECT_EQ(0u, q0_.Size());
}

// Tests Dequeue().
TEST_F(QueueTest, Dequeue) {
  int * n = q0_.Dequeue();
  EXPECT_TRUE(n == NULL);

  n = q1_.Dequeue();
  ASSERT_TRUE(n != NULL);
  EXPECT_EQ(1, *n);
  EXPECT_EQ(0u, q1_.Size());
  delete n;

  n = q2_.Dequeue();
  ASSERT_TRUE(n != NULL);
  EXPECT_EQ(2, *n);
  EXPECT_EQ(1u, q2_.Size());
  delete n;
}

// Tests the Queue::Map() function.
TEST_F(QueueTest, Map) {
  MapTester(&q0_);
  MapTester(&q1_);
  MapTester(&q2_);
}
