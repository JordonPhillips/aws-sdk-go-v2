package software.amazon.smithy.aws.go.codegen;

import java.util.Collections;
import java.util.Map;
import java.util.function.Predicate;
import java.util.logging.Logger;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.DocumentShapeDeserVisitor;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.shapes.CollectionShape;
import software.amazon.smithy.model.shapes.DocumentShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.utils.FunctionalUtils;

/**
 * Visitor to generate deserialization functions for shapes in XML protocol
 * document bodies.
 *
 * This class handles function body generation for all types expected by the
 * {@code DocumentShapeDeserVisitor}. No other shape type serialization is overwritten.
 */
public class XmlShapeDeserVisitor extends DocumentShapeDeserVisitor {
    private static final Logger LOGGER = Logger.getLogger(JsonShapeDeserVisitor.class.getName());
    private final Predicate<MemberShape> memberFilter;

    /**
     * @param context The generation context.
     */
    public XmlShapeDeserVisitor(GenerationContext context) {
        this(context, FunctionalUtils.alwaysTrue());
    }

    /**
     * @param context The generation context.
     * @param memberFilter A filter that is applied to structure members. This is useful for
     *     members that won't be in the body.
     */
    public XmlShapeDeserVisitor(GenerationContext context, Predicate<MemberShape> memberFilter) {
        super(context);
        this.memberFilter = memberFilter;
    }

    @Override
    protected Map<String, String> getAdditionalArguments() {
        return Collections.singletonMap("decoder", "*xml.Decoder");
    }

    @Override
    protected void deserializeCollection(GenerationContext context, CollectionShape shape) {
        context.getWriter().write("return nil");
    }

    @Override
    protected void deserializeDocument(GenerationContext context, DocumentShape shape) {
        GoWriter writer = context.getWriter();
        LOGGER.warning("The document type is unsupported for XML protocols.");
        writer.addUseImports(SmithyGoDependency.SMITHY);
        writer.write("return &smithy.DeserializationError{Err: fmt.Errorf("
                + "\"Document type is unsupported for XML protocols.\")}");
    }

    @Override
    protected void deserializeMap(GenerationContext context, MapShape shape) {
        context.getWriter().write("return nil");
    }

    @Override
    protected void deserializeStructure(GenerationContext context, StructureShape shape) {
        context.getWriter().write("return nil");
    }

    @Override
    protected void deserializeUnion(GenerationContext context, UnionShape shape) {
        GoWriter writer = context.getWriter();
        // TODO: implement union deserialization
        // The tricky part is going to be accumulating bytes for unknown members.
        LOGGER.warning("Union type is currently unsupported for XML deserialization.");
        context.getWriter().writeDocs("TODO: implement union serialization.");
        writer.write("return nil");
    }

}
