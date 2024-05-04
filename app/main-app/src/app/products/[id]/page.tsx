import Image from 'next/image';
import productImage from '../../../../public/images/product_detail.png';
import Link from 'next/link';
import {
  arrowLeftIcon,
  heartOutlineIcon,
  starIcon,
} from '../../../../public/icons';
import { ExpandableText } from '../../../../components/products/expand-text';
import ProductItem from '../../../../components/products/product-item';

export default function ProductDetailPage() {
  const descriptionText =
    'Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sint ipsum dolore officia aspernatur earum, maxime mollitia possimus enim, autem, quibusdam doloremque? Ut vel inventore itaque dignissimos eius ipsa, quae fuga!';
  const items = [
    {
      key: 'Size',
      prices: [0, 10, 100],
      properties: ['S', 'M', 'L'],
      require: true,
    },
    {
      key: 'Ice',
      prices: [0, 10, 100],
      properties: ['1/3', '1/2', '2'],
      require: false,
    },
    {
      key: 'Sugar',
      prices: [0, 10, 100],
      properties: ['1/3', '1/2', '2'],
      require: false,
    },
  ];

  return (
    <section>
      <div className="flex justify-between items-center">
        <div>
          <Link href="/">
            <Image
              width={24}
              height={24}
              src={arrowLeftIcon}
              alt="arrow_icon"
            />
          </Link>
        </div>
        <h2 className="text-xl font-medium">Detail</h2>
        <div>
          <Image
            width={24}
            height={24}
            src={heartOutlineIcon}
            alt="favorites_icon"
          />
        </div>
      </div>
      <div className="mt-[40px]">
        <Image className="w-full" src={productImage} alt="product_detail" />
      </div>
      <div className="flex items-center justify-between mt-[20px]">
        <div>
          <h3 className="font-semibold text-[1.6rem]">Cappuccino</h3>
        </div>
        <div className="flex items-center">
          <Image height={24} width={24} src={starIcon} alt="star_icon" />
          <p className="font-medium">4.8&nbsp;</p>
          <span className="text-xs text-[#808080]">(259)</span>
        </div>
      </div>
      <div className="w-full mx-0 my-3.5 border-[thin] border-solid border-[#EAEAEA]"></div>
      <div>
        <h4 className="font-semibold text-[1.2rem] mx-0 my-4">Description</h4>
        <ExpandableText>{descriptionText}</ExpandableText>
      </div>
      <div>
        {items.map((item, i) => (
          <ProductItem
            key={i}
            title={item.key}
            properties={item.properties}
            require={item.require}
          />
        ))}
      </div>
      <div className="flex justify-between items-center mx-0 my-5">
        <div>
          <p className="text-[#808080] text-sm">Price</p>
          <p className="text-[color:var(--primary-color)] font-semibold text-lg">
            $10.00
          </p>
        </div>
        <div className="w-[64%]">
          <button className="bg-[color:var(--primary-color)] text-white w-full px-5 py-2.5 rounded-[10px]">
            Add to Cart
          </button>
        </div>
      </div>
    </section>
  );
}
