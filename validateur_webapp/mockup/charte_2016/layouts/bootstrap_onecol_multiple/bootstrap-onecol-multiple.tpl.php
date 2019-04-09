<div class="<?php print $classes ?>" <?php if (!empty($css_id)) { print "id=\"$css_id\""; } ?>>
  <?php if ($content['navigation']): ?>
    <div class="panel-navigation">
      <div class="container">
        <div class="row row-navigation">
          <?php print $content['navigation']; ?>
        </div>
      </div>
    </div>
  <?php endif; ?>
  
  <?php if ($content['top'] || $content['middle'] || $content['bottom']): ?>
    <div class="panel-main">
      <?php if ($content['top']): ?>
        <div class="row row-top">
          <?php print $content['top']; ?>
        </div>
      <?php endif; ?>
      <?php if ($content['middle']): ?>
        <div class="row row-middle">
          <?php print $content['middle']; ?>
        </div>
      <?php endif; ?>
      <?php if ($content['bottom']): ?>
        <div class="row row-bottom">
          <?php print $content['bottom']; ?>
        </div>
      <?php endif; ?>
    </div>
  <?php endif; ?>
</div>