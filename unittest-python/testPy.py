import unittest

import calc

class CalTest(unittest.TestCase):
    def test_add_num_and_double(self):
        cal = calc.Cal()
        self.assertEqual(cal.add_num_and_double(1, 1), 4)

    def test_add_num_and_double_raise(self):
        cal = calc.Cal()
        # エラー処理ではwithを使う
        with self.assertRaises(ValueError):
            cal.add_num_and_double('1', '1')

if __name__ == "__main__":
    unittest.main()