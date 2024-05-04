interface Props {
  title: string;
  properties: {
    price: number;
    key: string;
  }[];
}

export default function ProductItem(props: Props) {
  return (
    <div>
      <h4>{props.title}</h4>
      <div>
        {props.properties.map((item, i) => (
          <button key={i}>{item.key}</button>
        ))}
      </div>
    </div>
  );
}
